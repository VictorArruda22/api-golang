package service

import (
	"github.com/VictorArruda22/api-golang/internal/dto"
	"github.com/VictorArruda22/api-golang/internal/repository"
	"github.com/VictorArruda22/api-golang/internal/utils"
)

type CategoryService struct {
	rp repository.CategoryRepository
}

func NewCategoryService(rp repository.CategoryRepository) CategoryService {
	return CategoryService{rp: rp}
}

func (s *CategoryService) CreateCategory(newCategory dto.CategoryRequestDTO) (dto.CategoryResponseDTO, error) {
	categoryEntitie := utils.ConvertCategoryDTOToEntitie(newCategory)
	newID, err := s.rp.CreateCategory(categoryEntitie)
	categoryResposeDTO := dto.CategoryResponseDTO{
		ID:          newID,
		Name:        categoryEntitie.Name,
		Description: categoryEntitie.Description,
	}
	return categoryResposeDTO, err
}
