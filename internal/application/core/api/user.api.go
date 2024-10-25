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

func (ua *UserApi) UpdateUser(id string, user domain.UpdateUserDTO) (*domain.User, error) {
	existingUser, err := ua.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	if user.Age != nil {
		existingUser.Age = *user.Age
	}
	if user.Name != nil {
		existingUser.Name = *user.Name
	}
	if user.Hobbies != nil {
		existingUser.Hobbies = *user.Hobbies
	}

	return ua.UserDbPort.Update(id, *existingUser)
}

func (ua *UserApi) DeleteUser(id string) error {
	return ua.UserDbPort.Delete(id)
}
