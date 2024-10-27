package services

import (
	"mereb_assessment/models"
	"mereb_assessment/repositories"
)

type PersonService struct {
	repo *repositories.PersonRepository
}

func NewPersonService(repo *repositories.PersonRepository) *PersonService {
	return &PersonService{repo: repo}
}

func (s *PersonService) GetAllPeople() ([]models.Person, error) {
	return s.repo.GetAll()
}

func (s *PersonService) GetPersonByID(id string) (models.Person, error) {
	return s.repo.GetByID(id)
}

func (s *PersonService) CreatePerson(person models.Person) (models.Person, error) {
	return s.repo.Create(person)
}

func (s *PersonService) UpdatePerson(id string, person models.Person) (models.Person, error) {
	return s.repo.Update(id, person)
}

func (s *PersonService) DeletePerson(id string) error {
	return s.repo.Delete(id)
}