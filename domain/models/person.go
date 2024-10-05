package models

import (
	"fmt"

	"github.com/beka-birhanu/GoCrudChallange/domain/common"
	errdmn "github.com/beka-birhanu/GoCrudChallange/domain/error"
	"github.com/google/uuid"
)

// Constants for name length validation.
const (
	minNameLength = 4
	maxNameLength = 100
)

// Person represents an individual with an ID, name, age, and hobbies.
type Person struct {
	id      uuid.UUID
	name    string
	age     int8
	hobbies []string
}

// PersonConfig is used to initialize a new Person.
type PersonConfig struct {
	Name    string
	Age     int8
	Hobbies []string
}

// NewPerson creates a new Person, validating the name and age, and initializing hobbies.
func NewPerson(c *PersonConfig) (*Person, ierr.IErr) {
	createdPerson := &Person{
		id: uuid.New(),
	}

	if err := createdPerson.SetName(c.Name); err != nil {
		return nil, err
	}

	if err := createdPerson.SetAge(c.Age); err != nil {
		return nil, err
	}

	createdPerson.SetHobbies(c.Hobbies)

	return createdPerson, nil
}

// Copy creates a deep copy of the Person struct.
func (p *Person) Copy() *Person {
	hobbiesCopy := make([]string, len(p.hobbies))
	copy(hobbiesCopy, p.hobbies)

	return &Person{
		id:      p.id,
		name:    p.name,
		age:     p.age,
		hobbies: hobbiesCopy,
	}
}

// ID returns the UUID of the person.
func (p *Person) ID() uuid.UUID {
	return p.id
}

// Name returns the name of the person.
func (p *Person) Name() string {
	return p.name
}

// SetName sets the person's name, ensuring it meets length requirements.
func (p *Person) SetName(newName string) ierr.IErr {
	if len(newName) < minNameLength || len(newName) > maxNameLength {
		return errdmn.NewValidation(fmt.Sprintf("Name must be in the range of [%d, %d]", minNameLength, maxNameLength))
	}
	p.name = newName
	return nil
}

// Age returns the person's age.
func (p *Person) Age() int8 {
	return p.age
}

// SetAge sets the person's age, ensuring it's non-negative.
func (p *Person) SetAge(newAge int8) ierr.IErr {
	if newAge < 0 {
		return errdmn.NewValidation("Age cannot be negative")
	}
	p.age = newAge
	return nil
}

// Hobbies returns the person's hobbies.
func (p *Person) Hobbies() []string {
	return p.hobbies
}

// SetHobbies replaces the existing hobbies with a new list, omitting empty strings.
func (p *Person) SetHobbies(newHobbies []string) {
	filteredHobbies := make([]string, 0)
	for _, hobby := range newHobbies {
		if hobby != "" { // Omit empty strings
			filteredHobbies = append(filteredHobbies, hobby)
		}
	}
	p.hobbies = filteredHobbies
}
