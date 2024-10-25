package dbport

import "github.com/Mahider-T/GoCrudChallange/internal/application/core/domain"

type PersonDbPort interface {
	Create(person domain.CreatePersonDTO) (*domain.Person, error)
	ReadOne(id string) (*domain.Person, error)
	ReadAll() ([]*domain.Person, error)
	Update(id string, person domain.Person) (*domain.Person, error)
	Delete(id string) error
}
