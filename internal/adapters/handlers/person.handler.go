package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Mahider-T/GoCrudChallange/internal/application/core/domain"
	"github.com/Mahider-T/GoCrudChallange/internal/ports/apiport"
)

type PersonHandler struct {
	PersonApi apiport.PersonApiPort
}

func NewPersonHandler(personApi apiport.PersonApiPort) *PersonHandler {
	return &PersonHandler{
		PersonApi: personApi,
	}
}

func (uh *PersonHandler) PersonCreateHandler(w http.ResponseWriter, r *http.Request) {
	var person domain.CreatePersonDTO
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdPerson, err := uh.PersonApi.CreatePerson(person)
	if err != nil {
		if errors.Is(err, domain.ErrDuplicateId) {
			http.Error(w, "Server failed to generate a unique id \nCould not create person", http.StatusInternalServerError)
			return
		} else if errors.Is(err, domain.ErrBadRequest) {
			http.Error(w, "Validation Failed.  \nCould not create person", http.StatusBadRequest)
			return
		}
		http.Error(w, "Could not create person", http.StatusInternalServerError)
		return
	}

	err = WriteJSON(w, http.StatusCreated, createdPerson, nil)
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}

}

func (uh *PersonHandler) PersonGetAllHandler(w http.ResponseWriter, r *http.Request) {

	persons, err := uh.PersonApi.GetAllPersons()

	if err != nil {
		if errors.Is(err, domain.ErrNoRecord) {
			http.Error(w, `{"error": "persons not found"}`, http.StatusNotFound)
			return
		}
		http.Error(w, "Could not get persons", http.StatusInternalServerError)
	}

	err = WriteJSON(w, http.StatusOK, persons, nil)
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}
}

func (uh *PersonHandler) PersonGetByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("personId")
	person, err := uh.PersonApi.GetPersonByID(id)
	if err != nil {
		if errors.Is(err, domain.ErrNoRecord) {
			http.Error(w, `{"error": "person not found"}`, http.StatusNotFound)
			return
		}
		http.Error(w, "Could not get person by id", http.StatusInternalServerError)
		return
	}

	err = WriteJSON(w, http.StatusOK, person, nil)
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}
}

func (uh *PersonHandler) PersonUpdateHandler(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("personId")

	var person domain.UpdatePersonDTO
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, "Invalid request payload \nCould not create person", http.StatusBadRequest)
		return
	} else if errors.Is(err, domain.ErrBadRequest) {
		http.Error(w, "Validation Failed. \nCould not create person", http.StatusBadRequest)
		return
	}

	updatedPerson, err := uh.PersonApi.UpdatePerson(id, person)

	if err != nil {
		if errors.Is(err, domain.ErrNoRecord) {
			http.Error(w, `{"error": "person not found"}`, http.StatusNotFound)
			return
		}
		http.Error(w, "Could not update person", http.StatusBadRequest)
		return
	}

	err = WriteJSON(w, http.StatusOK, updatedPerson, nil)
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}
}

func (uh *PersonHandler) PersonDeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("personId")
	err := uh.PersonApi.DeletePerson(id)

	if err != nil {
		if errors.Is(err, domain.ErrNoRecord) {
			http.Error(w, `{"error": "person not found"}`, http.StatusNotFound)
			return
		}
		http.Error(w, "Could not delete person", http.StatusInternalServerError)
		return
	}

	err = WriteJSON(w, http.StatusNoContent, "Successfully deleted person", nil) //task : returns 200 instead of 204
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}

}
