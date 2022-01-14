package repositories

import (
	"api/src/models"
	"database/sql"
)

//Publicacoes representa um repositorio de publicacoes
type Publicacoes struct {
	db *sql.DB
}

//PublicacaoRepository cria um repositorio de publicacoes
func PublicacaoRepository(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

//Create cria uma publicao no banco de dados
func (PublicationRepository Publicacoes) Create(publicao models.Publicacao) (uint64, error) {
	statement, err := PublicationRepository.db.Prepare("INSERT INTO publicacoes (titulo, conteudo, autor_id) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(publicao.Titulo, publicao.Conteudo, publicao.AutorId)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastID), nil
}

//BuscaPorID traz uma unica publicacao do banco de dados
func (PublicationRepository Publicacoes) BuscaPorID(publiId uint64) (models.Publicacao, error) {
	row, err := PublicationRepository.db.Query("SELECT p.*, u.nick from publicacoes p INNER JOIN usuarios u ON u.id = p.autor_id WHERE p.id = ?", publiId)
	if err != nil {
		return models.Publicacao{}, err
	}
	defer row.Close()

	var publication models.Publicacao

	if row.Next() {
		if err = row.Scan(
			&publication.ID,
			&publication.Titulo,
			&publication.Conteudo,
			&publication.AutorId,
			&publication.Curtidas,
			&publication.CriadoEm,
			&publication.AutorNick,
		); err != nil {
			return models.Publicacao{}, err
		}
	}

	return publication, nil
}

//Busca traz as publicacoes dos usuarios seguidos e tambem as propias publicacoes
func (PublicationRepository Publicacoes) Busca(userId uint64) ([]models.Publicacao, error) {
	rows, err := PublicationRepository.db.Query(`
		select distinct p.*, u.nick from publicacoes p 
		inner join usuarios u on u.id = p.autor_id 
		inner join seguidores s on p.autor_id = s.usuario_id 
		where u.id = ? or s.seguidor_id = ?
	`, userId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var publications []models.Publicacao

	for rows.Next() {
		var publication models.Publicacao

		if err = rows.Scan(
			&publication.ID,
			&publication.Titulo,
			&publication.Conteudo,
			&publication.AutorId,
			&publication.Curtidas,
			&publication.CriadoEm,
			&publication.AutorNick,
		); err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil

}
