package db

import (
	"errors"

	"github.com/Mahider-T/GoCrudChallange/internal/application/core/domain"
	"github.com/google/uuid"
)

type UserRepo struct {
	store map[string]domain.User
}

func NewUserRepo(store map[string]domain.User) *UserRepo {
	return &UserRepo{
		store: store,
	}
}
func (ur *UserRepo) Create(user domain.User) (*domain.User, error) {
	id := uuid.New().String()
	user.Id = id
	ur.store[id] = user
	// usr := ur.store[id] //Task
	usr, err := ur.ReadOne(id)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (ur *UserRepo) ReadOne(id string) (*domain.User, error) {
	usr := ur.store[id]
	return &usr, nil
}
func (ur *UserRepo) ReadAll() ([]*domain.User, error) {
	usrs := []*domain.User{}

	for _, value := range ur.store {
		usrs = append(usrs, &value)
	}
	return usrs, nil
}

func (ur *UserRepo) Update(id string, user domain.User) (*domain.User, error) {
	_, err := ur.ReadOne(id)

	if err != nil {
		return nil, err
	}
	ur.store[id] = user
	usr := ur.store[id]
	return &usr, nil
}

func (ur *UserRepo) Delete(id string) error {
	_, ok := ur.store[id]

	if !ok {
		return errors.New("no such record")
	}
	delete(ur.store, id)
	return nil
}
