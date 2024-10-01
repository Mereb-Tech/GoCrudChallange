package domain

import (
	"github.com/google/uuid"
)

type Person struct {
	ID 			uuid.UUID 		`json:"id"`
	Name 		string			`json:"name"`			
	Age 		int16			`json:"age"`
	Hobbies 	[]string		`json:"hobbies"`
}

type NewPerson struct {
	Name 		string			`json:"name"`
	Age 		int16			`json:"age"`
	Hobbies 	[]string		`json:"hobbies"`
}

type PersonUseCase interface {
	GetAllPersons() (*[]Person, error)
	GetPersonById(string) (Person, error)
	Register(*NewPerson) error
	UpdatePerson(NewPerson, string) ([]Person, error)
	DeletePerson(string)
}

type PersonRepository interface {
	GetAllPersons() (*[]Person, error)
	GetPersonById(string) (Person, error)
	Register(*NewPerson) error
	UpdatePerson(NewPerson, string) ([]Person, error)
	DeletePerson(string)
}

var InMemory = []Person{}