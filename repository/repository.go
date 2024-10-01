package repository

import (
	"errors"
	"mereb_go/domain"

	"github.com/google/uuid"
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

func (pr *PersonRepository) GetPersonById(person_id string) (domain.Person, error) {
	id, err := uuid.Parse(person_id)
	if err != nil {
		return domain.Person{}, errors.New("invalid id")
	}

	for _, value := range *pr.InMemory {
		if id == value.ID {
			return value, nil
		}
	}

	return domain.Person{}, errors.New("person with the specified id not found")
}

func (pr *PersonRepository) Register(newPerson *domain.NewPerson) (error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	addedPerson := domain.Person{
		ID: id,
		Name: newPerson.Name,
		Age: newPerson.Age,
		Hobbies: newPerson.Hobbies,
	}

	*pr.InMemory = append(*pr.InMemory, addedPerson)
	return nil
}

func (pr *PersonRepository) UpdatePerson(updatedInfo domain.NewPerson, person_id string) ([]domain.Person, error) {
	id, err := uuid.Parse(person_id)
	if err != nil {
		return []domain.Person{}, errors.New("invalid user id")
	}

	for index, value := range *pr.InMemory {
		if value.ID == id {
			if updatedInfo.Name != "" {
                (*pr.InMemory)[index].Name = updatedInfo.Name
            }
            if updatedInfo.Age != 0 {
                (*pr.InMemory)[index].Age = updatedInfo.Age
            }
			if updatedInfo.Hobbies != nil {
				(*pr.InMemory)[index].Hobbies= updatedInfo.Hobbies
			}
            return *pr.InMemory, nil
		}
	}

	return []domain.Person{}, errors.New("person with the specified id not founr")
}
