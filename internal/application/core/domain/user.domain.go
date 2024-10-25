package domain

import (
	"github.com/go-playground/validator/v10"
)

type User struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Hobbies []string `json:"hobbies"`
}

type UpdateUserDTO struct {
	Name    *string   `json:"name" validate:"omitempty,min=1"`
	Age     *int      `json:"age" validate:"omitempty,gte=0,lte=130"`
	Hobbies *[]string `json:"hobbies" validate:"omitempty,dive,min=1"`
}

type CreateUserDTO struct {
	Name    string   `json:"name" validate:"required,min=1"`
	Age     int      `json:"age" validate:"required,gte=0,lte=130"`
	Hobbies []string `json:"hobbies" validate:"required,dive,min=1"`
}

func ValidateCreateUserDTO(createUserDTO CreateUserDTO) error {
	validate := validator.New()
	return validate.Struct(createUserDTO)
}

func ValidateUpdateUserDTO(updateUserDTO UpdateUserDTO) error {
	validate := validator.New()
	return validate.Struct(updateUserDTO)
}
