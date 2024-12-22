package repository

import (
	"database/sql"

	"github.com/VictorArruda22/api-golang/internal/entities"
	"github.com/VictorArruda22/api-golang/internal/utils"
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
		return 0, utils.ErrCategoryRepositoryInternalError
	}
	newID, err := result.LastInsertId()
	if err != nil {
		return 0, utils.ErrCategoryRepositoryInternalError
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
		return utils.ErrCategoryRepositoryInternalError
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return utils.ErrCategoryRepositoryInternalError
	}

	if rowsAffected == 0 {
		return utils.ErrCategoryRepositoryNotFound
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
			return entities.Category{}, utils.ErrCategoryRepositoryNotFound
		}
		return entities.Category{}, utils.ErrCategoryRepositoryNotFound
	}

	return categorySearched, nil
}

func (r *CategoryRepository) GetAll() ([]entities.Category, error) {
	sqlStatement := `
        SELECT id, name, description FROM Category;`

	rows, err := r.db.Query(sqlStatement)
	if err != nil {
		return nil, utils.ErrCategoryRepositoryInternalError
	}
	defer rows.Close()

	var categories []entities.Category

	for rows.Next() {
		var category entities.Category
		if err := rows.Scan(&category.ID, &category.Name, &category.Description); err != nil {
			return nil, utils.ErrCategoryRepositoryInternalError
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, utils.ErrCategoryRepositoryInternalError
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
		return entities.Category{}, utils.ErrCategoryRepositoryInternalError
	}

	var updatedCategory entities.Category
	selectStatement := `
        SELECT id, name, description
        FROM Category
        WHERE id = ?;`

	err = r.db.QueryRow(selectStatement, id).Scan(&updatedCategory.ID, &updatedCategory.Name, &updatedCategory.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return entities.Category{}, utils.ErrCategoryRepositoryNotFound
		}
		return entities.Category{}, utils.ErrCategoryRepositoryInternalError
	}

	return updatedCategory, nil
}
