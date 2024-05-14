CREATE DATABASE mydb;

CREATE TABLE albums (
	id SERIAL PRIMARY KEY,
    title VARCHAR(255),
    artist VARCHAR(255),
    price DECIMAL(10, 2)
    );