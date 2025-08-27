CREATE DATABASE `crud-book-go`; USE
    `crud-book-go`;
CREATE TABLE books(
    Id INT AUTO_INCREMENT PRIMARY KEY,
    Title VARCHAR(255) NOT NULL,
    Author VARCHAR(255) NOT NULL,
    YEAR INT NOT NULL
); CREATE DATABASE `crud-book-go-test`; USE
    `crud-book-go-test`;
CREATE TABLE books(
    Id INT AUTO_INCREMENT PRIMARY KEY,
    Title VARCHAR(255) NOT NULL,
    Author VARCHAR(255) NOT NULL,
    YEAR INT NOT NULL
);