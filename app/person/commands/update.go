package personcmd

import (
	"github.com/beka-birhanu/GoCrudChallange/app/common/cqrs"
	irepo "github.com/beka-birhanu/GoCrudChallange/app/common/interface/repo"
	ierr "github.com/beka-birhanu/GoCrudChallange/domain/common"
	"github.com/beka-birhanu/GoCrudChallange/domain/models"
	"github.com/google/uuid"
)

// UpdatePersonCommand represents the command to update an existing person.
type UpdatePersonCommand struct {
	ID      uuid.UUID // ID of the person to be updated
	Name    string    // New name for the person
	Age     int8      // New age for the person
	Hobbies []string  // New list of hobbies for the person
}

// UpdatePersonHandler is the command handler for updating an existing person.
type UpdatePersonHandler struct {
	repo irepo.IPerson
}

// Ensure UpdatePersonHandler implements the Handler interface for handling UpdatePersonCommand and returning a Person.
var _ cqrs.Handler[*UpdatePersonCommand, *models.Person] = &UpdatePersonHandler{}

// NewUpdatePersonHandler creates a new instance of UpdatePersonHandler.
func NewUpdatePersonHandler(repo irepo.IPerson) *UpdatePersonHandler {
	return &UpdatePersonHandler{repo: repo}
}

// Handle processes the UpdatePersonCommand and returns the updated person or an error.
func (h *UpdatePersonHandler) Handle(command *UpdatePersonCommand) (*models.Person, ierr.IErr) {
	existingPerson, err := h.repo.Get(command.ID)
	if err != nil {
		return nil, err
	}

	if err := existingPerson.SetName(command.Name); err != nil {
		return nil, err
	}

	if err := existingPerson.SetAge(command.Age); err != nil {
		return nil, err
	}

	existingPerson.SetHobbies(command.Hobbies)

	if err := h.repo.Save(existingPerson); err != nil {
		return nil, err
	}

	return existingPerson, nil
}
