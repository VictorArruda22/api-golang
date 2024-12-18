package service

import "github.com/VictorArruda22/api-golang/internal/entities"

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
	return s.rp.Update(author)
}

func (s *AuthorService) Delete(id int) error {
	return s.rp.Delete(id)
}
