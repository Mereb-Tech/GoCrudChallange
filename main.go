package main

import (
    "github.com/go-playground/validator/v10"
)

type Person struct {
    ID      string   `json:"uid" validate:"required"`
    Name    string   `json:"name" validate:"required,min=2,max=100"`
    Age     int      `json:"age" validate:"gte=0,lte=120"`
    Hobbies []string `json:"hobbies" validate:"dive,required,min=1,max=50"`
}

// Example usage
func main() {
    validate := validator.New()
    person := Person{
        ID:      "123e4567-e89b-12d3-a456-426614174000",
        Name:    "s",
        Age:     30,
    }

    err := validate.Struct(person)
    if err != nil {
        // Handle validation errors
        println(err.Error())
    } else {
        println("Validation passed!")
    }
}
