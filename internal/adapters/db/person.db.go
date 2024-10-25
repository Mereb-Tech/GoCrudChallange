package db

import (
	"github.com/Mahider-T/GoCrudChallange/internal/application/core/domain"
	"github.com/google/uuid"
)

type PersonRepo struct {
	store map[string]domain.Person
}

func NewPersonRepo(store map[string]domain.Person) *PersonRepo {
	return &PersonRepo{
		store: store,
	}
}
func (ur *PersonRepo) Create(person domain.CreatePersonDTO) (*domain.Person, error) {
	domainPerson := domain.Person{
		Id:      uuid.New().String(),
		Name:    person.Name,
		Age:     person.Age,
		Hobbies: person.Hobbies,
	}
	if _, exists := ur.store[domainPerson.Id]; exists {
		return nil, domain.ErrDuplicateId
	}
	ur.store[domainPerson.Id] = domainPerson

	usr := ur.store[domainPerson.Id]
	return &usr, nil
}

func (ur *PersonRepo) ReadOne(id string) (*domain.Person, error) {
	usr, ok := ur.store[id]
	if !ok {
		return nil, domain.ErrNoRecord
	}
	return &usr, nil
}
func (ur *PersonRepo) ReadAll() ([]*domain.Person, error) {
	usrs := []*domain.Person{}

	for _, value := range ur.store {
		usrs = append(usrs, &value)
	}
	if len(usrs) == 0 {
		return nil, domain.ErrNoRecord
	}
	return usrs, nil
}

func (ur *PersonRepo) Update(id string, person domain.Person) (*domain.Person, error) {
	_, err := ur.ReadOne(id)

	if err != nil {
		return nil, err
	}
	ur.store[id] = person
	usr := ur.store[id]
	return &usr, nil
}

func (ur *PersonRepo) Delete(id string) error {
	_, ok := ur.store[id]

	if !ok {
		return domain.ErrNoRecord
	}
	delete(ur.store, id)
	return nil
}
