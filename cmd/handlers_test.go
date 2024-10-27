package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Mahider-T/GoCrudChallange/internal/adapters/handlers"
	"github.com/Mahider-T/GoCrudChallange/internal/application/core/domain"
	"github.com/Mahider-T/GoCrudChallange/internal/assert"
)

// Mock for PersonApiPort interface
type MockPersonApi struct {
	CreatePersonFunc  func(person domain.CreatePersonDTO) (*domain.Person, error)
	GetPersonByIDFunc func(id string) (*domain.Person, error)
	GetAllPersonsFunc func() ([]*domain.Person, error)
	UpdatePersonFunc  func(id string, person domain.UpdatePersonDTO) (*domain.Person, error)
	DeletePersonFunc  func(id string) error
}

func (m *MockPersonApi) CreatePerson(person domain.CreatePersonDTO) (*domain.Person, error) {
	return m.CreatePersonFunc(person)
}

func (m *MockPersonApi) GetPersonByID(id string) (*domain.Person, error) {
	return m.GetPersonByIDFunc(id)
}

func (m *MockPersonApi) GetAllPersons() ([]*domain.Person, error) {
	return m.GetAllPersonsFunc()
}

func (m *MockPersonApi) UpdatePerson(id string, person domain.UpdatePersonDTO) (*domain.Person, error) {
	return m.UpdatePersonFunc(id, person)
}

func (m *MockPersonApi) DeletePerson(id string) error {
	return m.DeletePersonFunc(id)
}

func TestPersonHandlers(t *testing.T) {
	mockPersonApi := &MockPersonApi{}
	personHandler := handlers.NewPersonHandler(mockPersonApi)

	t.Run("TestPersonGetAllHandler_Success", func(t *testing.T) {
		mockPersonApi.GetAllPersonsFunc = func() ([]*domain.Person, error) {
			return []*domain.Person{
				{Id: "1", Name: "Mahider Tekola", Age: 23},
				{Id: "2", Name: "Jane Tekola", Age: 23},
			}, nil
		}

		req := httptest.NewRequest("GET", "/person", nil)
		rr := httptest.NewRecorder()

		personHandler.PersonGetAllHandler(rr, req)

		assert.Equal(t, rr.Code, http.StatusOK)

		var persons []*domain.Person
		err := json.NewDecoder(rr.Body).Decode(&persons)
		assert.NilError(t, err)
		assert.Equal(t, len(persons), 2)
	})

	t.Run("TestPersonGetAllHandler_NoRecords", func(t *testing.T) {
		mockPersonApi.GetAllPersonsFunc = func() ([]*domain.Person, error) {
			return nil, domain.ErrNoRecord
		}

		req := httptest.NewRequest("GET", "/person", nil)
		rr := httptest.NewRecorder()

		personHandler.PersonGetAllHandler(rr, req)

		assert.Equal(t, rr.Code, http.StatusNotFound)
		assert.StringContains(t, rr.Body.String(), "persons not found")
	})

	t.Run("TestPersonGetByIdHandler_Success", func(t *testing.T) {
		personDTO := domain.CreatePersonDTO{
			Name: "Mahider Tekola",
			Age:  30,
		}

		mockPersonApi.CreatePersonFunc = func(person domain.CreatePersonDTO) (*domain.Person, error) {
			return &domain.Person{Name: person.Name, Age: person.Age}, nil
		}

		createdPerson, err := mockPersonApi.CreatePerson(personDTO)
		assert.NilError(t, err)

		assert.Equal(t, createdPerson.Id, createdPerson.Id)

		mockPersonApi.GetPersonByIDFunc = func(id string) (*domain.Person, error) {
			if id == createdPerson.Id {
				return &domain.Person{Id: createdPerson.Id, Name: "Mahider Tekola", Age: 23}, nil
			}
			return nil, domain.ErrNoRecord
		}

		req := httptest.NewRequest("GET", "/person/"+createdPerson.Id, nil)
		rr := httptest.NewRecorder()

		personHandler.PersonGetByIdHandler(rr, req)

		assert.Equal(t, rr.Code, http.StatusOK)

		var person domain.Person
		err = json.NewDecoder(rr.Body).Decode(&person)
		assert.NilError(t, err)
		assert.Equal(t, person.Id, createdPerson.Id)
		assert.Equal(t, person.Name, "Mahider Tekola")
		assert.Equal(t, person.Age, 30)
	})

	t.Run("TestPersonGetByIdHandler_NotFound", func(t *testing.T) {
		id := "non-existing-id"
		mockPersonApi.GetPersonByIDFunc = func(id string) (*domain.Person, error) {
			return nil, domain.ErrNoRecord
		}

		req := httptest.NewRequest("GET", "/person/"+id, nil)
		rr := httptest.NewRecorder()

		personHandler.PersonGetByIdHandler(rr, req)

		assert.Equal(t, rr.Code, http.StatusNotFound)
		assert.StringContains(t, rr.Body.String(), "person not found")
	})

	t.Run("TestPersonUpdateHandler_Success", func(t *testing.T) {

		personDTO := domain.CreatePersonDTO{
			Name:    "Mahider Tekola",
			Age:     23,
			Hobbies: []string{"reading", "traveling"},
		}

		mockPersonApi.CreatePersonFunc = func(person domain.CreatePersonDTO) (*domain.Person, error) {
			return &domain.Person{Name: person.Name, Age: person.Age}, nil
		}

		createdPerson, err := mockPersonApi.CreatePerson(personDTO)
		assert.NilError(t, err)

		updateDTO := domain.UpdatePersonDTO{
			Name:    stringPtr("Jane Tekola"),
			Age:     intPtr(23),
			Hobbies: &[]string{"reading", "traveling"},
		}

		mockPersonApi.UpdatePersonFunc = func(id string, person domain.UpdatePersonDTO) (*domain.Person, error) {

			if id != createdPerson.Id {
				t.Errorf("expected id %s, got %s", createdPerson.Id, id)
			}
			return &domain.Person{Id: createdPerson.Id, Name: *person.Name, Age: *person.Age}, nil
		}

		reqBody, _ := json.Marshal(updateDTO)
		req := httptest.NewRequest("PUT", "/person/"+createdPerson.Id, bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		personHandler.PersonUpdateHandler(rr, req)

		assert.Equal(t, rr.Code, http.StatusOK)

		var updatedPerson domain.Person
		err = json.NewDecoder(rr.Body).Decode(&updatedPerson)
		assert.NilError(t, err)
		assert.Equal(t, updatedPerson.Id, createdPerson.Id)
	})

	t.Run("TestPersonUpdateHandler_InvalidJSON", func(t *testing.T) {
		id := "existing-id"
		req := httptest.NewRequest("PUT", "/person/"+id, bytes.NewBufferString("invalid json"))
		rr := httptest.NewRecorder()

		personHandler.PersonUpdateHandler(rr, req)

		assert.Equal(t, rr.Code, http.StatusBadRequest)
		assert.StringContains(t, rr.Body.String(), "Invalid request payload")
	})

	t.Run("TestPersonDeleteHandler_Success", func(t *testing.T) {
		id := "existing-id"
		mockPersonApi.DeletePersonFunc = func(id string) error {
			return nil
		}

		req := httptest.NewRequest("DELETE", "/person/"+id, nil)
		rr := httptest.NewRecorder()

		personHandler.PersonDeleteHandler(rr, req)

		assert.Equal(t, rr.Code, http.StatusNoContent)
	})

	t.Run("TestPersonDeleteHandler_NotFound", func(t *testing.T) {
		id := "non-existing-id"
		mockPersonApi.DeletePersonFunc = func(id string) error {
			return domain.ErrNoRecord
		}

		req := httptest.NewRequest("DELETE", "/person/"+id, nil)
		rr := httptest.NewRecorder()

		personHandler.PersonDeleteHandler(rr, req)

		assert.Equal(t, rr.Code, http.StatusNotFound)
		assert.StringContains(t, rr.Body.String(), "person not found")
	})
}

// Helper functions to create pointers for testing
func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}
