package repository

import (
	"fmt"
	Domain "mereb/Domain"
	Models "mereb/Domain/Models"
)

type PersonRepository struct {
	persons map[string]Models.Person
}

func NewPersonRepository() Domain.PersonRepository {
	return &PersonRepository{
		persons: map[string]Models.Person{},
	}
}

func (repo *PersonRepository) GetAllPersons() ([]Models.Person, error) {
	persons := []Models.Person{}
	for _, person := range repo.persons {
		persons = append(persons, person)
	}
	if len(persons) == 0 {
		return nil, fmt.Errorf("no people found")
	}
	return persons, nil
}

func (repo *PersonRepository) GetPersonByID(id string) (Models.Person, error) {
	person := repo.persons[id]
	if person.ID == "" {
		return Models.Person{}, fmt.Errorf("person with the given id not found")
	}
	return person, nil
}

func (repo *PersonRepository) CreatePerson(person Models.Person) (Models.Person, error) {
	repo.persons[person.ID] = person
	return person, nil
}

func (repo *PersonRepository) UpdatePerson(id string, person Models.Person) (Models.Person, error) {
	person.ID = id
	repo.persons[id] = person
	return person, nil
}

func (repo *PersonRepository) DeletePerson(id string) error {
	delete(repo.persons, id)
	return nil
}
