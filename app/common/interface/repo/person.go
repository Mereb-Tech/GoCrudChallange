package irepo

import (
	ierr "github.com/beka-birhanu/GoCrudChallange/domain/common"
	"github.com/beka-birhanu/GoCrudChallange/domain/models"
	"github.com/google/uuid"
)

// IPerson defines the interface for managing Person entities in a repository.
type IPerson interface {
	// Save adds a new Person to the repository or updates an existing one.
	Save(person *models.Person) ierr.IErr

	// Get retrieves a Person by ID.
	Get(id uuid.UUID) (*models.Person, ierr.IErr)

	// Delete removes a Person from the repository.
	Delete(id uuid.UUID) ierr.IErr

	// GetAll retrieves all Persons from the repository.
	GetAll() ([]*models.Person, ierr.IErr)
}
