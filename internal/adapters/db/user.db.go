package db

import (
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
func (ur *UserRepo) Create(user domain.CreateUserDTO) (*domain.User, error) {
	domainUser := domain.User{
		Id:      uuid.New().String(),
		Name:    user.Name,
		Age:     user.Age,
		Hobbies: user.Hobbies,
	}
	if _, exists := ur.store[domainUser.Id]; exists {
		return nil, domain.ErrDuplicateId
	}
	ur.store[domainUser.Id] = domainUser

	usr := ur.store[domainUser.Id]
	return &usr, nil
}

func (ur *UserRepo) ReadOne(id string) (*domain.User, error) {
	usr, ok := ur.store[id]
	if !ok {
		return nil, domain.ErrNoRecord
	}
	return &usr, nil
}
func (ur *UserRepo) ReadAll() ([]*domain.User, error) {
	usrs := []*domain.User{}

	for _, value := range ur.store {
		usrs = append(usrs, &value)
	}
	if len(usrs) == 0 {
		return nil, domain.ErrNoRecord
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
		return domain.ErrNoRecord
	}
	delete(ur.store, id)
	return nil
}
