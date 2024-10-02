package usecases

import (
	"github.com/abe16s/GoCrudChallange/domain"
	"github.com/google/uuid"
)

type IService interface {
	CreatePerson(name string, age int, hobbies []string) error
	GetAllPersons() ([]*domain.Person, error)
	GetPersonById(id uuid.UUID) (*domain.Person, error)
	UpdatePerson(id uuid.UUID, name string, age int, hobbies []string) error
	DeletePerson(id uuid.UUID) error
}

type Service struct {
	repo IRepo
}

func NewService(repo IRepo) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreatePerson(name string, age int, hobbies []string) error {
	person := domain.NewPerson(name, age, hobbies)
	return s.repo.Create(person)
}

func (s *Service) GetAllPersons() ([]*domain.Person, error) {
	return s.repo.GetAllPersons()
}

func (s *Service) GetPersonById(id uuid.UUID) (*domain.Person, error) {
	return s.repo.GetPersonById(id)
}

func (s *Service) UpdatePerson(id uuid.UUID, name string, age int, hobbies []string) error {
	person := domain.NewPerson(name, age, hobbies)
	return s.repo.Update(person)
}

func (s *Service) DeletePerson(id uuid.UUID) error {
	return s.repo.Delete(id)
}