package utils

import (
	"github.com/VictorArruda22/api-golang/internal/dto"
	"github.com/VictorArruda22/api-golang/internal/entities"
	"github.com/joho/godotenv"
)

func LoadEnv() error {
	err := godotenv.Load("../.env")
	if err != nil {
		return err
	}
	return nil
}

func ConvertCategoryDTOToEntitie(newCategory dto.CategoryRequestDTO) entities.Category {
	return entities.Category{
		Name:        newCategory.Name,
		Description: newCategory.Description,
	}
}
