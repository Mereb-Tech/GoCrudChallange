package tests

import (
	dtos "GoCrudChallenge/Domain/DTOs"
	models "GoCrudChallenge/Domain/Models"
	usecases "GoCrudChallenge/UseCases"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// MockPersonRepository provides a mock implementation of the PersonRepository interface
type MockPersonRepository struct {
	mock.Mock
}

func (m *MockPersonRepository) CreatePerson(person *models.Person) error {
	args := m.Called(person)
	return args.Error(0)
}

func (m *MockPersonRepository) GetPersonByID(id string) (*models.Person, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Person), args.Error(1)
}

func (m *MockPersonRepository) UpdatePerson(person *models.Person) error {
	args := m.Called(person)
	return args.Error(0)
}

func (m *MockPersonRepository) DeletePerson(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockPersonRepository) GetAllPersons() []*models.Person {
	args := m.Called()
	return args.Get(0).([]*models.Person)
}

// UseCaseTestSuite defines the test suite
type UseCaseTestSuite struct {
	suite.Suite
	useCase  *usecases.PersonUseCase
	mockRepo *MockPersonRepository
}

// SetupTest sets up the use case and mock repository before each test
func (suite *UseCaseTestSuite) SetupTest() {
	suite.mockRepo = new(MockPersonRepository)
	suite.useCase = &usecases.PersonUseCase{
		Repo:      suite.mockRepo,
		Validator: validator.New(),
	}
}

// TestCreatePerson tests creating a person
func (suite *UseCaseTestSuite) TestCreatePerson() {
	personRequest := dtos.PersonRequestDTO{
		Name:    "John Doe",
		Age:     30,
		Hobbies: []string{"Reading", "Hiking"},
	}

	person := &models.Person{
		ID:      uuid.NewString(),
		Name:    personRequest.Name,
		Age:     personRequest.Age,
		Hobbies: personRequest.Hobbies,
	}

	suite.mockRepo.On("CreatePerson", mock.AnythingOfType("*models.Person")).Return(nil)

	createdPerson, err := suite.useCase.CreatePerson(personRequest)
	suite.NoError(err, "CreatePerson should not return an error")
	suite.Equal(person.Name, createdPerson.Name, "Person's name should be 'John Doe'")
}

// TestGetPersonByID tests getting a person by ID
func (suite *UseCaseTestSuite) TestGetPersonByID() {
	person := &models.Person{
		ID:      uuid.New().String(),
		Name:    "Jane Doe",
		Age:     25,
		Hobbies: []string{"Cooking", "Cycling"},
	}
	suite.mockRepo.On("GetPersonByID", person.ID).Return(person, nil)

	fetchedPerson, err := suite.useCase.GetPersonByID(person.ID)
	suite.NoError(err, "GetPersonByID should not return an error")
	suite.Equal(person.ID, fetchedPerson.ID, "Fetched person should have the same ID")
}

// TestUpdatePerson tests updating a person
func (suite *UseCaseTestSuite) TestUpdatePerson() {
	personID := uuid.NewString()
	personRequest := dtos.PersonRequestDTO{
		Name:    "Alice Updated",
		Age:     29,
		Hobbies: []string{"Running"},
	}

	person := &models.Person{
		ID:      personID,
		Name:    "Alice Updated",
		Age:     29,
		Hobbies: []string{"Running"},
	}

	suite.mockRepo.On("UpdatePerson", person).Return(nil)
	suite.mockRepo.On("GetPersonByID", personID).Return(person, nil)

	updatedPerson, err := suite.useCase.UpdatePerson(personID, personRequest)
	suite.NoError(err, "UpdatePerson should not return an error")
	suite.Equal(person.Name, updatedPerson.Name, "Updated name should match")
}

// TestDeletePerson tests deleting a person
func (suite *UseCaseTestSuite) TestDeletePerson() {
	personID := uuid.New().String()
	suite.mockRepo.On("DeletePerson", personID).Return(nil)

	err := suite.useCase.DeletePerson(personID)
	suite.NoError(err, "DeletePerson should not return an error")
}

// TestGetAllPersons tests retrieving all persons
func (suite *UseCaseTestSuite) TestGetAllPersons() {
	person1 := &models.Person{
		ID:      uuid.New().String(),
		Name:    "Chris",
		Age:     34,
		Hobbies: []string{"Gaming", "Skiing"},
	}
	person2 := &models.Person{
		ID:      uuid.New().String(),
		Name:    "Dana",
		Age:     22,
		Hobbies: []string{"Reading", "Writing"},
	}

	persons := []*models.Person{person1, person2}
	suite.mockRepo.On("GetAllPersons").Return(persons)

	allPersons, err := suite.useCase.GetAllPersons()
	suite.NoError(err, "GetAllPersons should not return an error")
	suite.Len(allPersons, 2, "There should be two persons returned")
}

// Run the test suite
func TestUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UseCaseTestSuite))
}
