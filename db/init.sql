CREATE DATABASE IF NOT EXISTS bookstore;
USE bookstore;

-- Table Author
CREATE TABLE Author (
    Id INT PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(255) NOT NULL,
    BirthDate DATE NOT NULL,
    Nationality VARCHAR(255) NOT NULL
);

-- Table Publisher
CREATE TABLE Publisher (
    Id INT PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(255) NOT NULL,
    Country VARCHAR(255) NOT NULL,
    FoundationYear DATE NOT NULL
);

-- Table Category
CREATE TABLE Category (
    Id INT PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(255) NOT NULL,
    Description VARCHAR(255)
);

-- Table Customer
CREATE TABLE Customer (
    Id INT PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(255) NOT NULL,
    Email VARCHAR(255) NOT NULL UNIQUE,
    Phone VARCHAR(20) NOT NULL
);

-- Table Book
CREATE TABLE Book (
    Id INT PRIMARY KEY AUTO_INCREMENT,
    Title VARCHAR(255) NOT NULL,
    Author INT NOT NULL,
    ISBN VARCHAR(13) UNIQUE NOT NULL,
    PublicationDate DATE,
    Publisher INT NOT NULL,
    FOREIGN KEY (Author) REFERENCES Author(Id),
    FOREIGN KEY (Publisher) REFERENCES Publisher(Id)
);

-- Table Borrow
CREATE TABLE Borrow (
    Id INT PRIMARY KEY AUTO_INCREMENT,
    Customer INT NOT NULL,
    Book INT NOT NULL,
    StartDate DATE NOT NULL,
    EndDate DATE,
    FOREIGN KEY (Customer) REFERENCES Customer(Id),
    FOREIGN KEY (Book) REFERENCES Book(Id)
);