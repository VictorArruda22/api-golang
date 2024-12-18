package repository

import (
	"database/sql"

	"github.com/VictorArruda22/api-golang/internal/entities"
)

type authorRepositoryImpl struct {
	db *sql.DB
}

func NewAuthorRepository(db *sql.DB) entities.AuthorRepository {
	return &authorRepositoryImpl{db: db}
}

// Create implements entities.AuthorRepository.
func (r *authorRepositoryImpl) Create(author entities.Author) ([]entities.Author, error) {
	result, err := r.db.Exec("INSERT INTO Author (name, nascimento, nascionalidade) VALUES (?, ?, ?)",
		author.Name, author.Nascimento, author.Nacionalidade)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	author.ID = int(id)
	return []entities.Author{author}, nil
}

// Delete implements entities.AuthorRepository.
func (r *authorRepositoryImpl) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM author WHERE id = ?", id)
	return err
}

// GetAll implements entities.AuthorRepository.
func (r *authorRepositoryImpl) GetAll() ([]entities.Author, error) {
	rows, err := r.db.Query("SELECT id, name, nascimento, nacionalidade FROM author")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []entities.Author
	for rows.Next() {
		var author entities.Author
		if err := rows.Scan(&author.ID, &author.Name, &author.Nascimento, &author.Nacionalidade); err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}
	return authors, nil
}

// GetByID implements entities.AuthorRepository.
func (r *authorRepositoryImpl) GetByID(id int) ([]entities.Author, error) {
	var author entities.Author
	err := r.db.QueryRow("SELECT id, name, nascimento, nacionalidade FROM Author WHERE id = ?", id).
		Scan(&author.ID, &author.Name, &author.Nascimento, &author.Nacionalidade)
	if err != nil {
		return []entities.Author{author}, err
	}
	return []entities.Author{author}, nil
}

// Update implements entities.AuthorRepository.
func (r *authorRepositoryImpl) Update(author entities.Author) ([]entities.Author, error) {
	_, err := r.db.Exec("UPDATE Author SET name = ?, nascimento = ?, nacionalidade = ? WHERE id = ?",
		author.Name, author.Nascimento, author.Nacionalidade)
	return nil, err
}
