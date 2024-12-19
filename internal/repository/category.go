package repository

import (
	"database/sql"
	"errors"

	"github.com/VictorArruda22/api-golang/internal/entities"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return CategoryRepository{db: db}
}

func (r *CategoryRepository) CreateCategory(categoryAdd entities.Category) (int, error) {
	sqlStatement := `
		INSERT INTO Category (Name, Description)
		VALUES (?, ?);`

	result, err := r.db.Exec(sqlStatement, categoryAdd.Name, categoryAdd.Description)
	if err != nil {
		return 0, err
	}
	newID, err := result.LastInsertId()
	if err != nil {
		return 0, errors.New("Erro ao buscar último ID inserido")
	}
	newIDiNT := int(newID)
	return newIDiNT, nil
}
