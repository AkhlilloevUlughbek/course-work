--execute this when you enter like psql -U postgres
create user scientist with login password 'scientist';
create database researches;
grant all privileges on database researches to scientist;


CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    email VARCHAR(100) UNIQUE,
    password varchar(30) not null ,
    first_name varchar(30) not null ,
    last_name varchar(30) not null ,
    organization varchar(30) not null ,
    country varchar(30) not null ,
    status varchar(30) not null ,
    category varchar(30) not null
);
