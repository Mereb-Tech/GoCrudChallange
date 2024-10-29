package test

import (
	Domain "mereb/Domain"
	Models "mereb/Domain/Models"
	repository "mereb/Repository"
	"testing"

	"github.com/stretchr/testify/suite"
)

type repositorySuite struct {
	suite.Suite
	repository Domain.PersonRepository
}

func (suite *repositorySuite) SetupTest() {
	test_repository := repository.NewPersonRepository()
	suite.repository = test_repository
}

var person = Models.Person{
	ID:      "123",
	Name:    "John",
	Age:     25,
	Hobbies: []string{"reading", "swimming"},
}

func (suite *repositorySuite) TestCreatePersonPositive() {
	createdPerson, err := suite.repository.CreatePerson(person)
	suite.NoError(err)
	suite.Equal(person, createdPerson)
}

func (suite *repositorySuite) TestGetPersonPositive() {
	id := person.ID
	incomingPerson, err := suite.repository.GetPersonByID(id)
	suite.Nil(err, "Person should be found")
	suite.Equal(person.Age, incomingPerson.Age, "Incoming person and person should be same")
	suite.Equal(person.Name, incomingPerson.Name, "Incoming person and person should be same")
	suite.Equal(person.Hobbies[0], incomingPerson.Hobbies[0], "Incoming person and person should be same")
	suite.Equal(person.Hobbies[1], incomingPerson.Hobbies[1], "Incoming person and person should be same")

}

func (suite *repositorySuite) TestGetAllPersonsPositive() {
	persons, err := suite.repository.GetAllPersons()
	suite.Nil(err, "There should be no error")
	suite.Equal(persons[0], person, "Person found should be same as the person created")
}

func (suite *repositorySuite) TestUpdatePersonPositive() {
	id := person.ID
	updatedPerson := Models.Person{
		ID:      id, // Ensure ID is set for the update
		Name:    "Updated John",
		Age:     26,
		Hobbies: []string{"reading updated", "swimming updated"},
	}

	_, err := suite.repository.UpdatePerson(id, updatedPerson)
	suite.Nil(err, "There should be no error")
	suite.Equal(updatedPerson, updatedPerson, "Person should be updated")
}

func (suite *repositorySuite) TestDeletePersonPositive() {
	id := person.ID
	err := suite.repository.DeletePerson(id)
	suite.Nil(err, "There should be no error")
}

func (suite *repositorySuite) TestGetPersonNegative() {
	id := person.ID
	_, err := suite.repository.GetPersonByID(id)
	suite.NotNil(err, "There should be error")
	suite.Equal(Models.Person{}, Models.Person{}, "Person should be empty")
}

func (suite *repositorySuite) TestSequentialTests() {
	suite.Run("Test Person Pos", suite.TestCreatePersonPositive)
	suite.Run("Test Get Person Pos", suite.TestGetPersonPositive)
	suite.Run("Test Get All Persons Pos", suite.TestGetAllPersonsPositive)
	suite.Run("Test Update Person Pos", suite.TestUpdatePersonPositive)
	suite.Run("Test Delete Person Pos", suite.TestDeletePersonPositive)
	suite.Run("Test Get Person Neg", suite.TestGetPersonNegative)
}

// Entry point for the test suite
func TestRepositorySuite(t *testing.T) {
	suite := new(repositorySuite)
	suite.SetT(t)
	suite.SetupTest() // Call setup to initialize
	suite.TestSequentialTests()
}
