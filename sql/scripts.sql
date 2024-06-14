--execute this when you enter like psql -U postgres
create user scientist with login password 'scientist';
create database researches;
grant all privileges on database researches to scientist;


CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    email VARCHAR(100) UNIQUE,
    password varchar(100) not null ,
    first_name varchar(100) not null ,
    last_name varchar(100) not null ,
    organization varchar(100) not null ,
    country varchar(100) not null ,
    status varchar(100) not null ,
    category varchar(100) not null,
    activated bool default false
);


create table otps(
    otp_id serial primary key ,
    email varchar(100) unique ,
    otp bigint unique,
    created_at timestamp default current_timestamp
);

create table tokens(
    token_id serial primary key ,
    email varchar(100) unique ,
    token text unique ,
    created_at timestamp default current_timestamp
);

create table researches(
    research_id serial primary key ,
    title varchar(50),
    description varchar(200),
    location varchar(200),
    user_email varchar(30),
    category varchar(50),
    approved bool default false
);

drop table users;
drop table otps;
drop table tokens;
drop table researches;
