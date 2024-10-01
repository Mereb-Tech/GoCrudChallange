package domain

import (
	"github.com/google/uuid"
)

type Person struct {
	ID uuid.UUID
	name string
	age int16
	hobbies []string
}

type NewPerson struct {
	name string
	age string
	hobbies []string
}

type PersonUseCase interface {
	GetAllPersons() (*[]Person, error)
	Register(*NewPerson) (error)
	UpdatePerson(string)
	DeletePerson(string)
}

type PersonRepository interface {
	GetAllPersons() (*[]Person, error)
	Register(*NewPerson) (error)
	UpdatePerson(string)
	DeletePerson(string)
}

var InMemory = []Person{}