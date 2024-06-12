CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE,
    password varchar(30) not null ,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
