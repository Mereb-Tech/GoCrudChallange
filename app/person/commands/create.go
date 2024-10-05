package personcmd

import (
	"github.com/beka-birhanu/GoCrudChallange/app/common/cqrs"
	irepo "github.com/beka-birhanu/GoCrudChallange/app/common/interface/repo"
	ierr "github.com/beka-birhanu/GoCrudChallange/domain/common"
	"github.com/beka-birhanu/GoCrudChallange/domain/models"
)

// CreatePersonCommand represents the command to create a new person.
type CreatePersonCommand struct {
	Name    string   // Name of the person
	Age     int8     // Age of the person
	Hobbies []string // List of hobbies for the person
}

// CreatePersonHandler is the command handler for creating a new person.
type CreatePersonHandler struct {
	repo irepo.IPerson
}

// Ensure CreatePersonHandler implements the Handler interface for handling CreatePersonCommand and returning a Person.
var _ cqrs.Handler[*CreatePersonCommand, *models.Person] = &CreatePersonHandler{}

// NewCreatePersonHandler creates a new instance of CreatePersonHandler.
func NewCreatePersonHandler(repo irepo.IPerson) *CreatePersonHandler {
	return &CreatePersonHandler{repo: repo}
}

// Handle processes the CreatePersonCommand and returns the created person or an error.
func (h *CreatePersonHandler) Handle(command *CreatePersonCommand) (*models.Person, ierr.IErr) {
	person, err := models.NewPerson(&models.PersonConfig{
		Name:    command.Name,
		Age:     command.Age,
		Hobbies: command.Hobbies,
	})
	if err != nil {
		return nil, err
	}

	if err := h.repo.Save(person); err != nil {
		return nil, err
	}

	return person, nil
}
