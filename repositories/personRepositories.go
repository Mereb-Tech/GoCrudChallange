package repositories

import (
	"errors"
	"github.com/google/uuid"
	"mereb_assessment/models"
)

type PersonRepository struct {
	data map[string]models.Person
}

func NewPersonRepository() *PersonRepository {
	return &PersonRepository{data: make(map[string]models.Person)}
}

func (r *PersonRepository) GetAll() ([]models.Person, error) {
	people := make([]models.Person, 0, len(r.data))
	for _, person := range r.data {
		people = append(people, person)
	}
	return people, nil
}

func (r *PersonRepository) GetByID(id string) (models.Person, error) {
	person, exists := r.data[id]
	if !exists {
		return models.Person{}, errors.New("person not found")
	}
	return person, nil
}

func (r *PersonRepository) Create(person models.Person) (models.Person, error) {
	person.ID = uuid.New().String()
	r.data[person.ID] = person
	return person, nil
}

func (r *PersonRepository) Update(id string, person models.Person) (models.Person, error) {
	if _, exists := r.data[id]; !exists {
		return models.Person{}, errors.New("person not found")
	}
	person.ID = id 
	r.data[id] = person
	return person, nil
}

func (r *PersonRepository) Delete(id string) error {
	if _, exists := r.data[id]; !exists {
		return errors.New("person not found")
	}
	delete(r.data, id)
	return nil
}