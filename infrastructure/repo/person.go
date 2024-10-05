package repo

import (
	"sync"

	ierr "github.com/beka-birhanu/GoCrudChallange/domain/common"
	errdmn "github.com/beka-birhanu/GoCrudChallange/domain/error"
	"github.com/beka-birhanu/GoCrudChallange/domain/models"
	"github.com/google/uuid"
)

// PersonRepo provides methods for managing Person entities.
type PersonRepo struct {
	mu     sync.RWMutex
	people map[uuid.UUID]*models.Person
}

// NewPersonRepo creates a new PersonRepo.
func NewPersonRepo() *PersonRepo {
	return &PersonRepo{
		people: make(map[uuid.UUID]*models.Person),
	}
}

// Save adds a new Person to the repository or updates an existing one.
func (r *PersonRepo) Save(person *models.Person) ierr.IErr {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.people[person.ID()] = person.Copy()
	return nil
}

// Get retrieves a Person by ID.
func (r *PersonRepo) Get(id uuid.UUID) (*models.Person, ierr.IErr) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	person, exists := r.people[id]
	if !exists {
		return nil, errdmn.NewNotFound("Person not found")
	}
	return person.Copy(), nil
}

// Delete removes a Person from the repository.
func (r *PersonRepo) Delete(id uuid.UUID) ierr.IErr {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.people[id]; !exists {
		return errdmn.NewNotFound("Person not found")
	}

	delete(r.people, id)
	return nil
}

// GetAll retrieves all Persons from the repository.
func (r *PersonRepo) GetAll() ([]*models.Person, ierr.IErr) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	peopleList := make([]*models.Person, 0, len(r.people))
	for _, person := range r.people {
		peopleList = append(peopleList, person.Copy())
	}
	return peopleList, nil
}
