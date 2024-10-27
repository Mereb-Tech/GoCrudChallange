package domain

import (
	"github.com/go-playground/validator/v10"
)

type Person struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Hobbies []string `json:"hobbies"`
}

type UpdatePersonDTO struct {
	Name    *string   `json:"name" validate:"omitempty,min=1"`
	Age     *int      `json:"age" validate:"omitempty,gte=0,lte=130"`
	Hobbies *[]string `json:"hobbies" validate:"omitempty,dive,min=1"`
}

type CreatePersonDTO struct {
	Name    string   `json:"name" validate:"required,min=1"`
	Age     int      `json:"age" validate:"required,gte=0,lte=130"`
	Hobbies []string `json:"hobbies" validate:"required,dive,min=1"`
}

func ValidateCreatePersonDTO(createPersonDTO CreatePersonDTO) error {
	validate := validator.New()
	return validate.Struct(createPersonDTO)
}

func ValidateUpdatePersonDTO(updatePersonDTO UpdatePersonDTO) error {
	validate := validator.New()
	return validate.Struct(updatePersonDTO)
}
