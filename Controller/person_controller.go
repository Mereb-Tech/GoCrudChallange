package controller

import (
	dtos "GoCrudChallenge/Domain/DTOs"
	interfaces "GoCrudChallenge/Domain/Interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PersonController struct {
	PersonUsecase interfaces.PersonUseCaseInterface
}

func NewPersonController(personUsecase interfaces.PersonUseCaseInterface) interfaces.PersonControllerInterface {
	return &PersonController{
		PersonUsecase: personUsecase,
	}
}

func (controller *PersonController) CreatePerson(c *gin.Context) {
	var personRequest dtos.PersonRequestDTO
	if err := c.ShouldBindJSON(&personRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	person, err := controller.PersonUsecase.CreatePerson(personRequest)
	if err != nil {
		if err.Error() == "Invalid person data" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid person data"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create person"})
		return
	}

	c.JSON(http.StatusCreated, person)
}

func (controller *PersonController) GetPersonByID(c *gin.Context) {
	id := c.Param("id")
	person, err := controller.PersonUsecase.GetPersonByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}

	c.JSON(http.StatusOK, person)
}

func (controller *PersonController) UpdatePerson(c *gin.Context) {
	id := c.Param("id")
	var personRequest dtos.PersonRequestDTO
	if err := c.ShouldBindJSON(&personRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	person, err := controller.PersonUsecase.UpdatePerson(id, personRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update person"})
		return
	}

	c.JSON(http.StatusOK, person)
}

func (controller *PersonController) DeletePerson(c *gin.Context) {
	id := c.Param("id")
	err := controller.PersonUsecase.DeletePerson(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete person"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (controller *PersonController) GetAllPersons(c *gin.Context) {
	persons, err := controller.PersonUsecase.GetAllPersons()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No persons found"})
		return
	}

	c.JSON(http.StatusOK, persons)
}
