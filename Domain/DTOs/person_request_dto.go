package dtos

type PersonRequestDTO struct {
	Name    string   `json:"name" validate:"required"`
	Age     int      `json:"age" validate:"required,gte=0"`
	Hobbies []string `json:"hobbies" validate:"required"`
}