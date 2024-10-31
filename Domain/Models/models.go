package models

type Person struct {
	ID      string   `json:"uid" validate:"required"`
	Name    string   `json:"name" validate:"required,min=2,max=100, pattern=^[a-zA-Z\s-]+$"`
	Age     int      `json:"age" validate:"required,gte=1,lte=150"`
	Hobbies []string `json:"hobbies" validate:"required,dive,required,min=2,max=50,pattern=^[a-zA-Z0-9\s-]+$"`

}

