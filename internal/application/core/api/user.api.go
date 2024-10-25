package api

import (
	"github.com/Mahider-T/GoCrudChallange/internal/application/core/domain"
	"github.com/Mahider-T/GoCrudChallange/internal/ports/dbport"
)

type UserApi struct {
	UserDbPort dbport.UserDbPort
}

func NewUserApi(userDbPort dbport.UserDbPort) *UserApi {
	return &UserApi{
		UserDbPort: userDbPort,
	}
}

func (ua *UserApi) CreateUser(user domain.User) (*domain.User, error) {
	return ua.UserDbPort.Create(user)

}

func (ua *UserApi) GetUserByID(id string) (*domain.User, error) {
	return ua.UserDbPort.ReadOne(id)
}

func (ua *UserApi) GetAllUsers() ([]*domain.User, error) {
	return ua.UserDbPort.ReadAll()
}

func (ua *UserApi) UpdateUser(id string, user domain.User) (*domain.User, error) {
	return ua.UserDbPort.Update(id, user)
}
