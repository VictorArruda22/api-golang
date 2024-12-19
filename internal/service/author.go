package service

import (
	"errors"

	"github.com/VictorArruda22/api-golang/internal/entities"
)

type AuthorService struct {
	rp entities.AuthorRepository
}

func NewAuthorService(rp entities.AuthorRepository) *AuthorService {
	return &AuthorService{rp: rp}
}

func (s *AuthorService) GetAll() ([]entities.Author, error) {
	return s.rp.GetAll()
}

func (s *AuthorService) GetByID(id int) ([]entities.Author, error) {
	return s.rp.GetByID(id)
}

func (s *AuthorService) Create(author entities.Author) ([]entities.Author, error) {
	return s.rp.Create(author)
}

func (s *AuthorService) Update(author entities.Author) ([]entities.Author, error) {
	// Verifique se o autor existe
	existingAuthor, err := s.rp.GetByID(author.ID)
	if err != nil {
		return nil, err
	}
	if len(existingAuthor) == 0 {
		return nil, errors.New("autor não encontrado")
	}

	return s.rp.Update(author)
}

func (s *AuthorService) Delete(id int) error {
	// Verifique se o autor existe
	existingAuthor, err := s.rp.GetByID(id)
	if err != nil {
		return err
	}
	if len(existingAuthor) == 0 {
		return errors.New("autor não encontrado")
	}
	return s.rp.Delete(id)
}
