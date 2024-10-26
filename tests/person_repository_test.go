package tests

import (
	models "GoCrudChallenge/Domain/Models"
	repository "GoCrudChallenge/Repository"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

// RepositoryTestSuite defines the test suite
type RepositoryTestSuite struct {
	suite.Suite
	repo *repository.PersonRepository
}

// SetupTest sets up the repository before each test
func (suite *RepositoryTestSuite) SetupTest() {
	suite.repo = repository.NewPersonRepository()
}

// TestCreatePerson tests the creation of a person
func (suite *RepositoryTestSuite) TestCreatePerson() {
	person := &models.Person{
		ID:      uuid.New().String(),
		Name:    "John Doe",
		Age:     30,
		Hobbies: []string{"Reading", "Hiking"},
	}

	err := suite.repo.CreatePerson(person)
	suite.NoError(err, "CreatePerson should not return an error")

	allPersons := suite.repo.GetAllPersons()
	suite.Len(allPersons, 1, "There should be one person in the repository")
	suite.Equal("John Doe", allPersons[0].Name, "Person's name should be 'John Doe'")
}

// TestGetPersonByID tests fetching a person by ID
func (suite *RepositoryTestSuite) TestGetPersonByID() {
	person := &models.Person{
		ID:      uuid.New().String(),
		Name:    "Jane Doe",
		Age:     25,
		Hobbies: []string{"Cooking", "Cycling"},
	}
	suite.repo.CreatePerson(person)

	fetchedPerson, err := suite.repo.GetPersonByID(person.ID)
	suite.NoError(err, "GetPersonByID should not return an error")
	suite.Equal(person.ID, fetchedPerson.ID, "Fetched person should have the same ID")
	suite.Equal(person.Name, fetchedPerson.Name, "Fetched person should have the correct name")
}

// TestUpdatePerson tests updating a person
func (suite *RepositoryTestSuite) TestUpdatePerson() {
	person := &models.Person{
		ID:      uuid.New().String(),
		Name:    "Alice",
		Age:     28,
		Hobbies: []string{"Painting", "Dancing"},
	}
	suite.repo.CreatePerson(person)

	updatedPerson := &models.Person{
		ID:      person.ID,
		Name:    "Alice Updated",
		Age:     29,
		Hobbies: []string{"Running"},
	}

	err := suite.repo.UpdatePerson(updatedPerson)
	suite.NoError(err, "UpdatePerson should not return an error")

	fetchedPerson, err := suite.repo.GetPersonByID(person.ID)
	suite.NoError(err, "GetPersonByID should not return an error")
	suite.Equal("Alice Updated", fetchedPerson.Name, "Name should be updated")
	suite.Equal(29, fetchedPerson.Age, "Age should be updated")
	suite.Equal([]string{"Running"}, fetchedPerson.Hobbies, "Hobbies should be updated")
}

// TestDeletePerson tests deleting a person
func (suite *RepositoryTestSuite) TestDeletePerson() {
	person := &models.Person{
		ID:      uuid.New().String(),
		Name:    "Bob",
		Age:     40,
		Hobbies: []string{"Gardening", "Swimming"},
	}
	suite.repo.CreatePerson(person)

	err := suite.repo.DeletePerson(person.ID)
	suite.NoError(err, "DeletePerson should not return an error")

	_, err = suite.repo.GetPersonByID(person.ID)
	suite.Error(err, "GetPersonByID should return an error for a deleted person")

	allPersons := suite.repo.GetAllPersons()
	suite.Len(allPersons, 0, "There should be no persons in the repository after deletion")
}

// TestGetAllPersons tests retrieving all persons
func (suite *RepositoryTestSuite) TestGetAllPersons() {
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

	suite.repo.CreatePerson(person1)
	suite.repo.CreatePerson(person2)

	allPersons := suite.repo.GetAllPersons()
	suite.Len(allPersons, 2, "There should be two persons in the repository")
	suite.Equal("Chris", allPersons[0].Name, "First person's name should be 'Chris'")
	suite.Equal("Dana", allPersons[1].Name, "Second person's name should be 'Dana'")
}

// Run the test suite
func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}
