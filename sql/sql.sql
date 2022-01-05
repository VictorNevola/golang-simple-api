CREATE DATABASE IF NOT EXISTS `devbook`;
USE `devbook`;

DROP TABLE IF EXISTS `usuarios`;

CREATE TABLE `usuarios` (
  id int auto_increment primary key,
  nome varchar(50) NOT NULL,
  nick varchar(50) NOT NULL unique,
  email varchar(50) NOT NULL unique,
  senha varchar(20) NOT NULL unique,
  criadoEm timestamp default current_timestamp(),
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
