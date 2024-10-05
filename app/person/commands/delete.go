package personcmd

import (
	"github.com/beka-birhanu/GoCrudChallange/app/common/cqrs"
	irepo "github.com/beka-birhanu/GoCrudChallange/app/common/interface/repo"
	ierr "github.com/beka-birhanu/GoCrudChallange/domain/common"
	"github.com/google/uuid"
)

// DeletePersonHandler is the command handler for deleting an existing person.
type DeletePersonHandler struct {
	repo irepo.IPerson
}

// Ensure DeletePersonHandler implements the Handler interface for handling DeletePersonCommand and returning a success status.
var _ cqrs.Handler[uuid.UUID, bool] = &DeletePersonHandler{}

// NewDeletePersonHandler creates a new instance of DeletePersonHandler.
func NewDeletePersonHandler(repo irepo.IPerson) *DeletePersonHandler {
	return &DeletePersonHandler{repo: repo}
}

// Handle processes the DeletePersonCommand and returns a success status and an error.
func (h *DeletePersonHandler) Handle(id uuid.UUID) (bool, ierr.IErr) {
	if err := h.repo.Delete(id); err != nil {
		return false, err
	}
	return true, nil
}
