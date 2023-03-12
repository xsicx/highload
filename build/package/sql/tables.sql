create table users (
    id UUID primary key,
    first_name VARCHAR(100) NOT NULL,
    second_name VARCHAR(100) NOT NULL,
    birthdate DATE NOT NULL,
    biography TEXT NULL,
    city VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL
);