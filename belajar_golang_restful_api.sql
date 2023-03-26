SHOW DATABASES;

CREATE DATABASE belajar_golang_restful_api;
USE belajar_golang_restful_api;

CREATE TABLE category (
	id INTEGER PRIMARY KEY AUTO_INCREMENT,
	name VARCHAR(200) NOT NULL,
	PRIMARY KEY(id)
)ENGINE=InnoDB;

SHOW TABLES;
DESC category;
SELECT * FROM category;

SHOW DATABASES;

CREATE DATABASE belajar_golang_restful_api_test;

USE belajar_golang_restful_api_test;

CREATE TABLE category (
	id INTEGER PRIMARY KEY AUTO_INCREMENT,
	name VARCHAR(200) NOT NULL
)ENGINE=InnoDB;

SHOW TABLES;

SELECT * FROM category;	