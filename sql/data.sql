INSERT INTO usuarios (nome, nick, email, senha)
values 
('João', 'joao', 'joao@gmail.com', '$2a$10$Wkyoxk90Q9dQiBcoLnBRhOZvG1ecSV/u6MrhQhv1ZwSbtfduswrA.'),
('Maria', 'maria', 'maria@gmail.com', '$2a$10$Wkyoxk90Q9dQiBcoLnBRhOZvG1ecSV/u6MrhQhv1ZwSbtfduswrA.'),
('Pedro', 'pedro', 'pedro@gmail.com', '$2a$10$Wkyoxk90Q9dQiBcoLnBRhOZvG1ecSV/u6MrhQhv1ZwSbtfduswrA.');

INSERT into seguidores (usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);

INSERT into publicacoes (titulo, conteudo, autor_id)
values
("Publicacao do usuario 1", "Essa a publicação do usuario 1", 1),
("Publicacao do usuario 2", "Essa a publicação do usuario 2", 2),
("Publicacao do usuario 3", "Essa a publicação do usuario 3", 3);