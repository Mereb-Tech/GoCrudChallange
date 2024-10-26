package interfaces

import (
	"github.com/gin-gonic/gin"
)



type PersonControllerInterface interface {
	CreatePerson(c *gin.Context)
	GetPersonByID(c *gin.Context)
	UpdatePerson(c *gin.Context)
	DeletePerson(c *gin.Context)
	GetAllPersons(c *gin.Context)
}