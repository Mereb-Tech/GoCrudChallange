package domain 

import "github.com/google/uuid"

type Person struct {
	ID      uuid.UUID   `json:"id"`
    Name    string   	`json:"name" binding:"required"`
    Age     int      	`json:"age" binding:"required"`
    Hobbies []string 	`json:"hobbies" binding:"required"`
}

func NewPerson(name string, age int, hobbies []string) Person {
    return Person{
        ID:      uuid.New(),
        Name:    name,
        Age:     age,
        Hobbies: hobbies,
    }
}