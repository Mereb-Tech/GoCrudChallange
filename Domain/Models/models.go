package models

type Person struct {
	ID      string   `json:"uid" validate:"required"`
	Name    string   `json:"name" validate:"required,min=2,max=100"`
	Age     int      `json:"age" validate:"required,gte=1,lte=150"`
	Hobbies []string `json:"hobbies" validate:"dive,min=1,max=50"`
}
