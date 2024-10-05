package controller

import (
	"github.com/google/uuid"
)

// CreatePersonDTO is the request payload for creating a person.
type CreatePersonDTO struct {
	Name    string   `json:"name" binding:"required,min=4,max=100"` // Name must be between 4 and 100 characters long.
	Age     int8     `json:"age" binding:"required,min=0"`          // Age must be a non-negative integer.
	Hobbies []string `json:"hobbies"`                               // List of hobbies (optional).
}

// PersonResponseDTO is the response payload for a single person.
type PersonResponseDTO struct {
	ID      uuid.UUID `json:"id"`      // Unique identifier for the person.
	Name    string    `json:"name"`    // Name of the person.
	Age     int8      `json:"age"`     // Age of the person.
	Hobbies []string  `json:"hobbies"` // List of hobbies for the person.
}
