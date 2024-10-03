package repositories

import (
	"errors"

	"github.com/abe16s/GoCrudChallange/domain"
	"github.com/google/uuid"
)

type Repository struct {
	persons map[uuid.UUID]domain.Person
}

func NewRepository() *Repository {
	return &Repository{
		persons: make(map[uuid.UUID]domain.Person),
	}
}

func (r *Repository) Create(person domain.Person) error {
	r.persons[person.ID] = person
	return nil
}

func (r *Repository) GetAllPersons() ([]*domain.Person, error) {
	allPersons := make([]*domain.Person, 0, len(r.persons))
	for _, person := range r.persons {
		allPersons = append(allPersons, &person)
	}
	return allPersons, nil
}

func (r *Repository) GetPersonById(id uuid.UUID) (*domain.Person, error) {
	if person, ok := r.persons[id]; ok {
		return &person, nil
	}
	return nil, errors.New("person not found")
}

func (r *Repository) Update(id uuid.UUID, updatePerson domain.Person) error {
	if _, ok := r.persons[id]; ok {
		r.persons[id] = updatePerson
		return nil
	}
	return errors.New("person not found")
}

func (r *Repository) Delete(id uuid.UUID) error {
	if _, ok := r.persons[id]; ok {
		delete(r.persons, id)
		return nil
	}
	return errors.New("person not found")
}

