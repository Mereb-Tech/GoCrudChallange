package models

type Person struct {
    ID      string   `json:"id"`
    Name    string   `json:"name" validate:"required"`
    Age     int      `json:"age" validate:"required,gte=0"` 
    Hobbies []string `json:"hobbies" validate:"required"`
}
