package personqry

import (
	"github.com/beka-birhanu/GoCrudChallange/app/common/cqrs"
	irepo "github.com/beka-birhanu/GoCrudChallange/app/common/interface/repo"
	ierr "github.com/beka-birhanu/GoCrudChallange/domain/common"
	"github.com/beka-birhanu/GoCrudChallange/domain/models"
)

// Ensure GetAllPersonHandler implements the cqrs.Handler interface.
var _ cqrs.Handler[struct{}, []*models.Person] = &GetAllPersonHandler{}

// GetAllPersonHandler is the query handler for retrieving all persons.
type GetAllPersonHandler struct {
	repo irepo.IPerson
}

// NewGetAllPersonHandler creates a new instance of GetAllPersonHandler.
func NewGetAllPersonHandler(repo irepo.IPerson) *GetAllPersonHandler {
	return &GetAllPersonHandler{repo: repo}
}

// Handle processes the request to retrieve all persons and returns the list or an error.
func (h *GetAllPersonHandler) Handle(_ struct{}) ([]*models.Person, ierr.IErr) {
	persons, err := h.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return persons, nil
}
