package services

import (
	"GoCrudChallange_Bisrat/data"
	"errors"
)

func GetAllPersons() ([]data.Person, error) {
    return data.Persons, nil
}

func GetPersonByID(id string) (data.Person, error) {
	for _, person := range data.Persons {
		if person.ID == id {
			return person, nil
		}
	}
	return data.Person{}, errors.New("person not found")
}

func CreatePerson(person *data.Person) error {
	data.Persons = append(data.Persons, *person)
	return nil
}

func UpdatePerson(id string, updatedPerson *data.Person) error {
	for i, person := range data.Persons {
		if person.ID == id {
			data.Persons[i] = *updatedPerson
			return nil
		}
	}
	return errors.New("person not found")
}

func DeletePerson(id string) error {
	for i, person := range data.Persons {
		if person.ID == id {
			data.Persons = append(data.Persons[:i], data.Persons[i+1:]...)
			return nil
		}
	}
	return errors.New("person not found")
}