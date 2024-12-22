package repository

import (
	"database/sql"
	"errors"

	"github.com/VictorArruda22/api-golang/internal/entities"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(categoryAdd entities.Category) (int, error) {
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

func (r *CategoryRepository) Delete(id int) error {
	sqlStatement := `
		DELETE FROM Category 
		WHERE id = ?;`

	result, err := r.db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("categoria nâo encontrada")
	}

	return nil
}

func (r *CategoryRepository) GetByID(id int) (entities.Category, error) {
	sqlStatement := `
        SELECT id, name, description FROM Category WHERE id = ?;`

	var categorySearched entities.Category
	err := r.db.QueryRow(sqlStatement, id).Scan(&categorySearched.ID, &categorySearched.Name, &categorySearched.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return entities.Category{}, errors.New("categoria não encontrada")
		}
		return entities.Category{}, err
	}

	return categorySearched, nil
}

func (r *CategoryRepository) GetAll() ([]entities.Category, error) {
	sqlStatement := `
        SELECT id, name, description FROM Category;`

	rows, err := r.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []entities.Category

	for rows.Next() {
		var category entities.Category
		if err := rows.Scan(&category.ID, &category.Name, &category.Description); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *CategoryRepository) Update(id int, categoryUpdated entities.Category) (entities.Category, error) {
	sqlStatement := `
        UPDATE Category
        SET name = ?, description = ?
        WHERE id = ?;`

	_, err := r.db.Exec(sqlStatement, categoryUpdated.Name, categoryUpdated.Description, id)
	if err != nil {
		return entities.Category{}, err
	}

	var updatedCategory entities.Category
	selectStatement := `
        SELECT id, name, description
        FROM Category
        WHERE id = ?;`

	err = r.db.QueryRow(selectStatement, id).Scan(&updatedCategory.ID, &updatedCategory.Name, &updatedCategory.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return entities.Category{}, errors.New("categoria não encontrada")
		}
		return entities.Category{}, err
	}

	return updatedCategory, nil
}
