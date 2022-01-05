package db

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Driver de conexao
)

//Connect é a conexao com o banco de dados
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.StringConnectionDB)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil

}
