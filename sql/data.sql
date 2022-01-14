INSERT INTO usuarios (nome, nick, email, senha)
values 
('João', 'joao', 'joao@gmail.com', '$2a$10$Wkyoxk90Q9dQiBcoLnBRhOZvG1ecSV/u6MrhQhv1ZwSbtfduswrA.'),
('Maria', 'maria', 'maria@gmail.com', '$2a$10$Wkyoxk90Q9dQiBcoLnBRhOZvG1ecSV/u6MrhQhv1ZwSbtfduswrA.'),
('Pedro', 'pedro', 'pedro@gmail.com', '$2a$10$Wkyoxk90Q9dQiBcoLnBRhOZvG1ecSV/u6MrhQhv1ZwSbtfduswrA.'),
('Ana', 'ana', 'ana@gmail.com', '$2a$10$Wkyoxk90Q9dQiBcoLnBRhOZvG1ecSV/u6MrhQhv1ZwSbtfduswrA.'),
('José', 'jose', 'jose@gmail.com', '$2a$10$Wkyoxk90Q9dQiBcoLnBRhOZvG1ecSV/u6MrhQhv1ZwSbtfduswrA.'),
('Joana', 'joana', 'joana@gmail.com', '$2a$10$Wkyoxk90Q9dQiBcoLnBRhOZvG1ecSV/u6MrhQhv1ZwSbtfduswrA.');

INSERT into seguidores (usuario_id, seguidor_id)
values
(1, 2),
(1, 3),
(1, 4),
(2, 1),
(2, 3),
(2, 4),
(3, 1),
(3, 2),
(3, 4),
(4, 1),
(4, 2),
(4, 3);

INSERT into publicacoes (titulo, conteudo, autor_id)
values
("Publicacao do usuario 1", "Essa a publicação do usuario 1", 1),
("Publicacao do usuario 2", "Essa a publicação do usuario 2", 2),
("Publicacao do usuario 3", "Essa a publicação do usuario 3", 3),
("Publicacao do usuario 4", "Essa a publicação do usuario 4", 4),
("Publicacao do usuario 5", "Essa a publicação do usuario 5", 5),
("Publicacao do usuario 6", "Essa a publicação do usuario 6", 6),
("Publicacao do usuario 7", "Essa a publicação do usuario 7", 7),
("Publicacao do usuario 8", "Essa a publicação do usuario 8", 8),
("Publicacao do usuario 9", "Essa a publicação do usuario 9", 9),
("Publicacao do usuario 10", "Essa a publicação do usuario 10", 10);