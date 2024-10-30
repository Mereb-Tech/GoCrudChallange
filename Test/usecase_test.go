package test

import (
	"mereb/Domain"
	models "mereb/Domain/Models"
	mocks "mereb/Test/Mocks"
	usecase "mereb/UseCase"
	"testing"

	"github.com/stretchr/testify/suite"
)

type usecaseSuite struct {
	suite.Suite
	usecase    Domain.PersonUsecase
	repository *mocks.PersonRepository
}

func (suite *usecaseSuite) SetupTestSuite() {
	// repository := mocks.NewPersonRepository(suite.T())
	suite.repository = new(mocks.PersonRepository)
	suite.usecase = usecase.NewPersonUsecase(suite.repository)
}

var createdPerson = models.Person{}
var person_uc = models.Person{
	Name:    "Jane Doe",
	Age:     25,
	Hobbies: []string{"Reading", "Swimming"},
}

func (suite *usecaseSuite) TestCreatePersonPositive() {
	// Mock the CreatePerson method

	suite.repository.On("CreatePerson", person_uc).Return(person_uc, nil)

	createdPerson, err := suite.usecase.CreatePerson(person_uc)
	suite.NoError(err, "error should be nil")
	suite.Equal(person_uc.Name, createdPerson.Name, "Name should be the same")
	suite.Equal(person_uc.Age, createdPerson.Age, "Age should be the same")
	suite.Equal(person_uc.Hobbies[0], createdPerson.Hobbies[0], "Hobbies should be the same")
}

func (suite *usecaseSuite) TestGetPersonPositive() {
	// Mock the GetPersonByID method
	id := createdPerson.ID
	incomingPerson, err := suite.usecase.GetPersonByID(id)
	suite.Nil(err, "Person should be found")
	suite.Equal(person_uc.Age, incomingPerson.Age, "Incoming person and person should be same")
	suite.Equal(person_uc.Name, incomingPerson.Name, "Incoming person and person should be same")
	suite.Equal(person_uc.Hobbies[0], incomingPerson.Hobbies[0], "Incoming person and person should be same")
	suite.Equal(person_uc.Hobbies[1], incomingPerson.Hobbies[1], "Incoming person and person should be same")
}

func (suite *usecaseSuite) TestUpdatePersonPositive() {
	// Mock the UpdatePerson method
	id := person_uc.ID
	newPerson := models.Person{
		Name:    "John Doe",
		Age:     30,
		Hobbies: []string{"Reading", "Swimming"},
	}
	updatedPerson, err := suite.usecase.UpdatePerson(id, newPerson)
	suite.NoError(err, "error should be nil")
	suite.Equal(newPerson.Name, updatedPerson.Name, "Name should be the same")
	suite.Equal(newPerson.Age, updatedPerson.Age, "Age should be the same")
	suite.Equal(newPerson.Hobbies[0], updatedPerson.Hobbies[0], "Hobbies should be the same")
	suite.Equal(newPerson.Hobbies[1], updatedPerson.Hobbies[1], "Hobbies should be the same")
}

func (suite *usecaseSuite) TestDeletePersonPositive() {
	// Mock the DeletePerson method
	id := createdPerson.ID
	err := suite.usecase.DeletePerson(id)
	suite.Nil(err, "Person should be deleted")
}

func (suite *usecaseSuite) TestSequentialUseCaseTests() {
	suite.Run("Test Person Pos", suite.TestCreatePersonPositive)
	// suite.Run("Test Get Person Pos", suite.TestGetPersonPositive)
	// suite.Run("Test Update Person Pos", suite.TestUpdatePersonPositive)
	// suite.Run("Test Delete Person Pos", suite.TestDeletePersonPositive)
}

// Entry point for the test suite
func TestUseCaseSuite(t *testing.T) {
	suite := new(usecaseSuite)
	suite.SetT(t)
	suite.SetupTestSuite() // Call setup to initialize
	suite.TestSequentialUseCaseTests()
}
