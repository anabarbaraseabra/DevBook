CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE users;

CREATE TABLE users (
    id varchar(16) primary key,
    name varchar(55) not null,
    nick varchar(55) not null unique,
    email varchar(55) not null unique,
    password varchar(55) not null,
    createdAt timestamp default current_timestamp()
) ENGINE=INNODB;