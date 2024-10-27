package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mereb_assessment/models"
	"mereb_assessment/services"
)

type PersonHandler struct {
	service *services.PersonService
}

func NewPersonHandler(service *services.PersonService) *PersonHandler {
	return &PersonHandler{service: service}
}

func (h *PersonHandler) RegisterRoutes(router *gin.Engine) {
	personRoutes := router.Group("/person")
	{
		personRoutes.GET("/", h.getAllPeople)
		personRoutes.GET("/:id", h.getPersonByID)
		personRoutes.POST("/", h.createPerson)
		personRoutes.PUT("/:id", h.updatePerson)
		personRoutes.DELETE("/:id", h.deletePerson)
	}
}

func (h *PersonHandler) getAllPeople(c *gin.Context) {
	people, err := h.service.GetAllPeople()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, people)
}

func (h *PersonHandler) getPersonByID(c *gin.Context) {
	id := c.Param("id")
	person, err := h.service.GetPersonByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, person)
}

func (h *PersonHandler) createPerson(c *gin.Context) {
	var person models.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	createdPerson, err := h.service.CreatePerson(person)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdPerson)
}

func (h *PersonHandler) updatePerson(c *gin.Context) {
	id := c.Param("id")
	var person models.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	updatedPerson, err := h.service.UpdatePerson(id, person)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedPerson)
}

func (h *PersonHandler) deletePerson(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeletePerson(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
