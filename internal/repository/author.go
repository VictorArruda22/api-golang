package repository

import (
	"database/sql"
	"errors"

	"github.com/VictorArruda22/api-golang/internal/entities"
	"github.com/VictorArruda22/api-golang/internal/utils"
	"github.com/go-sql-driver/mysql"
)

type authorRepositoryImpl struct {
	db *sql.DB
}

func NewAuthorRepository(db *sql.DB) entities.AuthorRepository {
	return &authorRepositoryImpl{db: db}
}

// Create implements entities.AuthorRepository.
func (r *authorRepositoryImpl) Create(author entities.Author) ([]entities.Author, error) {
	result, err := r.db.Exec("INSERT INTO Author (name, BirthDate, Nationality) VALUES (?, ?, ?)",
		author.Name, author.Nascimento, author.Nacionalidade)
	if err != nil {
		return nil, handlerError(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, handlerError(err)
	}
	author.ID = int(id)
	return []entities.Author{author}, nil
}

// Delete implements entities.AuthorRepository.
func (r *authorRepositoryImpl) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM Author WHERE id = ?", id)
	return handlerError(err)
}

// GetAll implements entities.AuthorRepository.
func (r *authorRepositoryImpl) GetAll() ([]entities.Author, error) {
	rows, err := r.db.Query("SELECT Id, Name, BirthDate, Nationality FROM Author")
	if err != nil {
		return nil, handlerError(err)
	}
	defer rows.Close()

	var authors []entities.Author
	for rows.Next() {
		var author entities.Author
		if err := rows.Scan(&author.ID, &author.Name, &author.Nascimento, &author.Nacionalidade); err != nil {
			return nil, handlerError(err)
		}
		authors = append(authors, author)
	}
	return authors, nil
}

// GetByID implements entities.AuthorRepository.
func (r *authorRepositoryImpl) GetByID(id int) ([]entities.Author, error) {
	var author entities.Author
	err := r.db.QueryRow("SELECT id, name, BirthDate, Nationality FROM Author WHERE id = ?", id).
		Scan(&author.ID, &author.Name, &author.Nascimento, &author.Nacionalidade)
	if err != nil {
		return nil, handlerError(err)
	}
	return []entities.Author{author}, nil
}

// Update implements entities.AuthorRepository.
func (r *authorRepositoryImpl) Update(author entities.Author) ([]entities.Author, error) {
	_, err := r.db.Exec("UPDATE Author SET name = ?, BirthDate = ?, Nationality = ? WHERE id = ?",
		author.Name, author.Nascimento, author.Nacionalidade, author.ID)
	return nil, handlerError(err)
}

func handlerError(err error) error {
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		switch mysqlErr.Number {
		case 1062:
			if mysqlErr.Message == "Duplicate entry" {
				return utils.ErrAuthorRepositoryInvalidID
			}
			err = utils.ErrAuthorRepositoryDuplicated
		case 1054:
			err = utils.ErrAuthorRepositoryBadField
		case 1048:
			err = utils.ErrAuthorRepositoryNullValue
		default:
			return utils.ErrAuthorRepositoryInternalError
		}
	}
	if errors.Is(err, sql.ErrNoRows) {
		return utils.ErrAuthorRepositoryNotFound
	}

	return err
}
