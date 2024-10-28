package repository

import (
	"fmt"
	Domain "mereb/Domain"
)

type PersonRepository struct {
	persons map[string]Domain.Person
}

func NewPersonRepository() Domain.PersonRepository {
	return &PersonRepository{
		persons: map[string]Domain.Person{},
	}
}

func (repo *PersonRepository) GetAllPersons() ([]Domain.Person, error) {
	persons := []Domain.Person{}
	for _, person := range repo.persons {
		persons = append(persons, person)
	}
	if len(persons) == 0 {
		return nil, fmt.Errorf("no people found")
	}
	return persons, nil

}
func (repo *PersonRepository) GetPersonByID(id string) (Domain.Person, error) {
	person := repo.persons[id]
	if person.ID == "" {
		return Domain.Person{}, fmt.Errorf("person with the given id not found")
	}
	return person, nil
}
func (repo *PersonRepository) CreatePerson(person Domain.Person) (Domain.Person, error) {
	repo.persons[person.ID] = person
	fmt.Print(person, repo.persons)
	return person, nil
}
func (repo *PersonRepository) UpdatePerson(id string, person Domain.Person) (Domain.Person, error) {
	person.ID = id
	repo.persons[id] = person
	return person, nil
}

func (repo *PersonRepository) DeletePerson(id string) ( error) {
	delete(repo.persons, id)
	return nil
}
