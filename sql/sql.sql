CREATE DATABASE IF NOT EXISTS `devbook`;
USE `devbook`;

DROP TABLE IF EXISTS `publicacoes`;
DROP TABLE IF EXISTS `seguidores`;
DROP TABLE IF EXISTS `usuarios`;

CREATE TABLE `usuarios` (
  id int auto_increment primary key,
  nome varchar(50) NOT NULL,
  nick varchar(50) NOT NULL unique,
  email varchar(50) NOT NULL unique,
  senha varchar(100) NOT NULL,
  criadoEm timestamp default current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `seguidores` (
  usuario_id int NOT NULL,
  FOREIGN KEY (usuario_id)
  REFERENCES usuarios(id)
  ON DELETE CASCADE,

  seguidor_id int NOT NULL,
  FOREIGN KEY (seguidor_id)
  REFERENCES usuarios(id)
  ON DELETE CASCADE,

  primary key (usuario_id, seguidor_id)

) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `publicacoes` (
  id int auto_increment primary key,
  titulo varchar(100) NOT NULL,
  conteudo varchar(300) NOT NULL,

  autor_id int NOT NULL,
  FOREIGN KEY (autor_id)
  REFERENCES usuarios(id)
  ON DELETE CASCADE,
  
  curtidas int default 0,
  criadoEm timestamp default current_timestamp()

) ENGINE=InnoDB DEFAULT CHARSET=utf8;
