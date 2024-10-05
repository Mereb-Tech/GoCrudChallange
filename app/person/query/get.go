package personqry

import (
	"github.com/beka-birhanu/GoCrudChallange/app/common/cqrs"
	irepo "github.com/beka-birhanu/GoCrudChallange/app/common/interface/repo"
	ierr "github.com/beka-birhanu/GoCrudChallange/domain/common"
	"github.com/beka-birhanu/GoCrudChallange/domain/models"
	"github.com/google/uuid"
)

// Ensure GetPersonHandler implements the cqrs.Handler interface.
var _ cqrs.Handler[uuid.UUID, *models.Person] = &GetPersonHandler{}

// GetPersonHandler is the query handler for retrieving a person by ID.
type GetPersonHandler struct {
	repo irepo.IPerson // Repository for managing Person entities
}

// NewGetPersonHandler creates a new instance of GetPersonHandler.
func NewGetPersonHandler(repo irepo.IPerson) *GetPersonHandler {
	return &GetPersonHandler{repo: repo}
}

// Handle processes the GetPersonQry and returns the requested person or an error.
func (h *GetPersonHandler) Handle(id uuid.UUID) (*models.Person, ierr.IErr) {
	person, err := h.repo.Get(id)
	if err != nil {
		return nil, err
	}

	return person, nil
}
