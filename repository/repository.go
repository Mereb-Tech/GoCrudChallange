package repository

import (
	"mereb_go/domain"
)

func NewPersonRepository(db *[]domain.Person) domain.PersonRepository {
	return &PersonRepository{
		InMemory: db,
	}
}


type PersonRepository struct {
	InMemory *[]domain.Person
}

func (pr *PersonRepository) DeletePerson(string) {
	panic("unimplemented")
}

func (pr *PersonRepository) GetAllPersons() (*[]domain.Person, error) {
	return pr.InMemory, nil
}

func (pr *PersonRepository) Register(*domain.NewPerson) (error) {
	panic("unimplemented")
}

func (pr *PersonRepository) UpdatePerson(string) {
	panic("unimplemented")
}
