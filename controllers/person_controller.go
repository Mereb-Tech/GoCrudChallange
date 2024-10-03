package controllers

import (
	"GoCrudChallange_Bisrat/data"
	"GoCrudChallange_Bisrat/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllPersons(c *gin.Context) {
	persons,err := services.GetAllPersons()
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	  c.JSON(http.StatusOK, persons)
}

func GetPerson(c *gin.Context) {
	id := c.Param("id")
    person, err := services.GetPersonByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
        return
    }
    c.JSON(http.StatusOK, person)

}

func CreatePerson(c *gin.Context) {
	 var person data.Person
    if err := c.ShouldBindJSON(&person); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := services.CreatePerson(&person)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, person)

}

func UpdatePerson(c *gin.Context) {
     id := c.Param("id")
    var updatedPerson data.Person
    if err := c.ShouldBindJSON(&updatedPerson); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := services.UpdatePerson(id, &updatedPerson)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
        return
    }
    c.JSON(http.StatusOK, updatedPerson)

}

func DeletePerson(c *gin.Context) {
    id := c.Param("id")
    err := services.DeletePerson(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message":"person deleted successfully"})

}