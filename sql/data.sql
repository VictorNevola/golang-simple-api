INSERT INTO usuarios (nome, nick, email, senha)
values 
('João', 'joao', 'joao@gmail.com', '$2a$10$ByMlQQpTT3LwwY8uNPT9pu1zENTtG/qcQyE3kj9h1eA0Qb4.IuYh2'),
('Maria', 'maria', 'maria@gmail.com', '$2a$10$ByMlQQpTT3LwwY8uNPT9pu1zENTtG/qcQyE3kj9h1eA0Qb4.IuYh2'),
('Pedro', 'pedro', 'pedro@gmail.com', '$2a$10$ByMlQQpTT3LwwY8uNPT9pu1zENTtG/qcQyE3kj9h1eA0Qb4.IuYh2'),
('Ana', 'ana', 'ana@gmail.com', '$2a$10$ByMlQQpTT3LwwY8uNPT9pu1zENTtG/qcQyE3kj9h1eA0Qb4.IuYh2'),
('José', 'jose', 'jose@gmail.com', '$2a$10$ByMlQQpTT3LwwY8uNPT9pu1zENTtG/qcQyE3kj9h1eA0Qb4.IuYh2'),
('Joana', 'joana', 'joana@gmail.com', '$2a$10$ByMlQQpTT3LwwY8uNPT9pu1zENTtG/qcQyE3kj9h1eA0Qb4.IuYh2');

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