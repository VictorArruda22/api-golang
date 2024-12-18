CREATE DATABASE IF NOT EXISTS bookstore;
USE bookstore;

-- Tabela Autor
CREATE TABLE Autor (
    Id INT PRIMARY KEY AUTO_INCREMENT,
    Nome VARCHAR(255) NOT NULL
    Nascimento DATE NOT NULL,
    Nacionalidade VARCHAR(255) NOT NULL
);

-- Tabela Editora
CREATE TABLE Editora (
    Id INT PRIMARY KEY AUTO_INCREMENT,
    Nome VARCHAR(255) NOT NULL,
    Pais VARCHAR(255) NOT NULL,
    AnoFundacao DATE NOT NULL
);

-- Tabela Categoria
CREATE TABLE Categoria (
    Id INT PRIMARY KEY AUTO_INCREMENT,
    Nome VARCHAR(255) NOT NULL,
    Descricao VARCHAR(255)
);

-- Tabela Cliente
CREATE TABLE Cliente (
    Id INT PRIMARY KEY AUTO_INCREMENT,
    Nome VARCHAR(255) NOT NULL,
    Email VARCHAR(255) NOT NULL UNIQUE,
    Telefone VARCHAR(20) NOT NULL
);

-- Tabela Livro
CREATE TABLE Livro (
    Id INT PRIMARY KEY AUTO_INCREMENT,
    Titulo VARCHAR(255) NOT NULL,
    Autor INT NOT NULL,
    ISBN VARCHAR(13) UNIQUE NOT NULL,
    DataPublicacao DATE,
    Editora INT NOT NULL,
    FOREIGN KEY (Autor) REFERENCES Autor(Id),
    FOREIGN KEY (Editora) REFERENCES Editora(Id)
);

-- Tabela Aluguel
CREATE TABLE Borrow (
    Id INT PRIMARY KEY AUTO_INCREMENT,
    Cliente INT NOT NULL,
    Livro INT NOT NULL,
    DataInicio DATE NOT NULL,
    DataFim DATE,
    FOREIGN KEY (Cliente) REFERENCES Cliente(Id),
    FOREIGN KEY (Livro) REFERENCES Livro(Id)
);