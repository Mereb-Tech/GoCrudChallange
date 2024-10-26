package tests

import (
	controller "GoCrudChallenge/Controller"
	dtos "GoCrudChallenge/Domain/DTOs"
	interfaces "GoCrudChallenge/Domain/Interfaces"
	models "GoCrudChallenge/Domain/Models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// MockPersonUsecase mocks the PersonUsecase interface
type MockPersonUsecase struct {
	mock.Mock
}

func (m *MockPersonUsecase) CreatePerson(personRequest dtos.PersonRequestDTO) (*models.Person, error) {
	args := m.Called(personRequest)
	return args.Get(0).(*models.Person), args.Error(1)
}

func (m *MockPersonUsecase) GetPersonByID(id string) (*models.Person, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Person), args.Error(1)
}

func (m *MockPersonUsecase) UpdatePerson(id string, personRequest dtos.PersonRequestDTO) (*models.Person, error) {
	args := m.Called(id, personRequest)
	return args.Get(0).(*models.Person), args.Error(1)
}

func (m *MockPersonUsecase) DeletePerson(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockPersonUsecase) GetAllPersons() ([]*models.Person, error) {
	args := m.Called()
	return args.Get(0).([]*models.Person), args.Error(1)
}

// ControllerTestSuite defines the test suite for the controller layer
type ControllerTestSuite struct {
	suite.Suite
	router           *gin.Engine
	mockUsecase      *MockPersonUsecase
	personController  interfaces.PersonControllerInterface
}

// SetupTest initializes the test suite
func (suite *ControllerTestSuite) SetupTest() {
	suite.mockUsecase = new(MockPersonUsecase)
	suite.personController = controller.NewPersonController(suite.mockUsecase)
	suite.router = gin.Default()

	// Define routes for each handler
	suite.router.POST("/person", suite.personController.CreatePerson)
	suite.router.GET("/person/:id", suite.personController.GetPersonByID)
	suite.router.PUT("/person/:id", suite.personController.UpdatePerson)
	suite.router.DELETE("/person/:id", suite.personController.DeletePerson)
	suite.router.GET("/person", suite.personController.GetAllPersons)
}

// TestCreatePerson tests the CreatePerson controller method
func (suite *ControllerTestSuite) TestCreatePerson() {
	personRequest := dtos.PersonRequestDTO{
		Name:    "John Doe",
		Age:     30,
		Hobbies: []string{"Reading", "Swimming"},
	}
	person := &models.Person{
		ID:      uuid.NewString(),
		Name:    personRequest.Name,
		Age:     personRequest.Age,
		Hobbies: personRequest.Hobbies,
	}

	suite.mockUsecase.On("CreatePerson", personRequest).Return(person, nil)

	// Mock HTTP request
	body, _ := json.Marshal(personRequest)
	req, _ := http.NewRequest(http.MethodPost, "/person", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusCreated, resp.Code, "HTTP status should be 201 Created")
	suite.mockUsecase.AssertCalled(suite.T(), "CreatePerson", personRequest)
}

// TestGetPersonByID tests the GetPersonByID controller method
func (suite *ControllerTestSuite) TestGetPersonByID() {
	personID := uuid.NewString()
	person := &models.Person{
		ID:      personID,
		Name:    "Jane Doe",
		Age:     25,
		Hobbies: []string{"Running"},
	}

	suite.mockUsecase.On("GetPersonByID", personID).Return(person, nil)

	req, _ := http.NewRequest(http.MethodGet, "/person/"+personID, nil)
	resp := httptest.NewRecorder()
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusOK, resp.Code, "HTTP status should be 200 OK")
	suite.mockUsecase.AssertCalled(suite.T(), "GetPersonByID", personID)
}

// TestUpdatePerson tests the UpdatePerson controller method
func (suite *ControllerTestSuite) TestUpdatePerson() {
	personID := uuid.NewString()
	personRequest := dtos.PersonRequestDTO{
		Name:    "Updated Name",
		Age:     28,
		Hobbies: []string{"Drawing"},
	}
	updatedPerson := &models.Person{
		ID:      personID,
		Name:    personRequest.Name,
		Age:     personRequest.Age,
		Hobbies: personRequest.Hobbies,
	}

	suite.mockUsecase.On("UpdatePerson", personID, personRequest).Return(updatedPerson, nil)

	body, _ := json.Marshal(personRequest)
	req, _ := http.NewRequest(http.MethodPut, "/person/"+personID, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusOK, resp.Code, "HTTP status should be 200 OK")
	suite.mockUsecase.AssertCalled(suite.T(), "UpdatePerson", personID, personRequest)
}

// TestDeletePerson tests the DeletePerson controller method
func (suite *ControllerTestSuite) TestDeletePerson() {
	personID := uuid.NewString()
	suite.mockUsecase.On("DeletePerson", personID).Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, "/person/"+personID, nil)
	resp := httptest.NewRecorder()
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusNoContent, resp.Code, "HTTP status should be 204 No Content")
	suite.mockUsecase.AssertCalled(suite.T(), "DeletePerson", personID)
}

// TestGetAllPersons tests the GetAllPersons controller method
func (suite *ControllerTestSuite) TestGetAllPersons() {
	person1 := &models.Person{ID: uuid.NewString(), Name: "Alice", Age: 31, Hobbies: []string{"Reading"}}
	person2 := &models.Person{ID: uuid.NewString(), Name: "Bob", Age: 27, Hobbies: []string{"Music"}}
	persons := []*models.Person{person1, person2}

	suite.mockUsecase.On("GetAllPersons").Return(persons, nil)

	req, _ := http.NewRequest(http.MethodGet, "/person", nil)
	resp := httptest.NewRecorder()
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusOK, resp.Code, "HTTP status should be 200 OK")
	suite.mockUsecase.AssertCalled(suite.T(), "GetAllPersons")
}

// Run the test suite
func TestControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ControllerTestSuite))
}
