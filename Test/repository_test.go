package test

import (
	Domain "mereb/Domain"
	Models "mereb/Domain/Models"
	repository "mereb/Repository"

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

func (suite *repositorySuite) TestCreatePersonNegative() {
	badId := Models.Person{
		Name:    "John",
		Age:     25,
		Hobbies: []string{"reading", "swimming"},
	}
	_, err := suite.repository.CreatePerson(badId)
	suite.NotNil(err, "There should be error returned")
}

func (suite *repositorySuite) TestGetPersonPositive() {
	id := person.ID

	incomingPerson, err := suite.repository.GetPersonByID(id)

	suite.Nil(err, "Person should be found")
	// check incomingPerson and person are same
	suite.Equal(person, incomingPerson, "Incoming person and person should be same")
}

func (suite *repositorySuite) TestGetAllPersonsPositive() {
	persons, err := suite.repository.GetAllPersons()
	suite.Nil(err, "There should be no error")
	suite.Equal(persons[0], person, "Person found should be same as the person created")
}

func (suite *repositorySuite) TestUpdatePersonPositive() {
	id := person.ID
	updatedPerson := Models.Person{
		Name:    "Updated John",
		Age:     26,
		Hobbies: []string{"reading updated", "swimming updated"},
	}

	person, err := suite.repository.UpdatePerson(id, updatedPerson)
	suite.Nil(err, "There should be no error")
	suite.Equal(person, updatedPerson, "Person should be updated")
}

func (suite *repositorySuite) TestDeletePersonPositive() {
	id := person.ID
	err := suite.repository.DeletePerson(id)
	suite.Nil(err, "There should be no error")
}

func (suite *repositorySuite) TestGetPersonNegative() {
	id := person.ID
	person, err := suite.repository.GetPersonByID(id)
	suite.NotNil(err, "There should be error")
	suite.Equal(person, Models.Person{}, "Person should be empty")
}

func (suite *repositorySuite) TestPersonRepository() {
	suite.Run("CreatePersonPositive", func() {
		suite.TestCreatePersonPositive()
	})

	suite.Run("CreatePersonNegative", func() {
		suite.TestCreatePersonNegative()
	})

	suite.Run("GetPersonPositive", func() {
		suite.TestGetPersonPositive()
	})
	suite.Run("GetAllPersonsPositive", func() {
		suite.TestGetAllPersonsPositive()
	})

	suite.Run("UpdatePersonPositive", func() {
		suite.TestUpdatePersonPositive()
	})

	suite.Run("DeletePersonPositive", func() {
		suite.TestDeletePersonPositive()
	})

	suite.Run("GetPersonNegative", func() {
		suite.TestGetPersonNegative()
	})

}
