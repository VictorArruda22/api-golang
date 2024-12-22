package service

import (
	"github.com/VictorArruda22/api-golang/internal/dto"
	"github.com/VictorArruda22/api-golang/internal/entities"
	"github.com/VictorArruda22/api-golang/internal/utils"
)

type CategoryService struct {
	rp entities.CategoryRepository
}

func NewCategoryService(rp entities.CategoryRepository) *CategoryService {
	return &CategoryService{rp: rp}
}

func (s *CategoryService) Create(newCategory dto.CategoryRequestDTO) (dto.CategoryResponseDTO, error) {
	categoryEntitie := utils.ConvertCategoryDTOToEntitie(newCategory)
	newID, err := s.rp.Create(categoryEntitie)
	categoryResposeDTO := dto.CategoryResponseDTO{
		ID:          newID,
		Name:        categoryEntitie.Name,
		Description: categoryEntitie.Description,
	}
	return categoryResposeDTO, err
}

func (s *CategoryService) Delete(id int) error {
	err := s.rp.Delete(id)
	return err
}

func (s *CategoryService) GetByID(id int) (dto.CategoryResponseDTO, error) {
	categoryEntitie, err := s.rp.GetByID(id)
	categoryResponseDTO := dto.CategoryResponseDTO{
		ID:          categoryEntitie.ID,
		Name:        categoryEntitie.Name,
		Description: categoryEntitie.Description,
	}
	return categoryResponseDTO, err
}

func (s *CategoryService) GetAll() ([]dto.CategoryResponseDTO, error) {
	var categoryResponseDTO []dto.CategoryResponseDTO
	categoryEntitie, err := s.rp.GetAll()
	for _, category := range categoryEntitie {
		categoryResponseDTO = append(categoryResponseDTO, dto.CategoryResponseDTO{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}
	return categoryResponseDTO, err
}

func (s *CategoryService) Update(id int, categoryUpdated dto.CategoryRequestDTO) (dto.CategoryResponseDTO, error) {
	categoryUpdatedEntitie := utils.ConvertCategoryDTOToEntitie(categoryUpdated)
	categoryUpdatedSearched, err := s.rp.Update(id, categoryUpdatedEntitie)
	categoryUpdatedResponseDTO := dto.CategoryResponseDTO{
		ID:          categoryUpdatedSearched.ID,
		Name:        categoryUpdatedSearched.Name,
		Description: categoryUpdatedSearched.Description,
	}
	return categoryUpdatedResponseDTO, err
}
