package entities

import (
	"github.com/VictorArruda22/api-golang/internal/dto"
)

type Category struct {
	ID          int
	Name        string
	Description string
}

type CategoryService interface {
	GetAll() ([]dto.CategoryResponseDTO, error)
	GetByID(id int) (dto.CategoryResponseDTO, error)
	Create(newCategory dto.CategoryRequestDTO) (dto.CategoryResponseDTO, error)
	Update(id int, categoryUpdated dto.CategoryRequestDTO) (dto.CategoryResponseDTO, error)
	Delete(id int) error
}
type CategoryRepository interface {
	GetAll() ([]Category, error)
	GetByID(id int) (Category, error)
	Create(categoryAdd Category) (int, error)
	Update(id int, categoryUpdated Category) (Category, error)
	Delete(id int) error
}
