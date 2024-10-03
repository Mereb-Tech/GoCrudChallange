package controllers

import (
	"GoCrudChallange_Bisrat/data"
	"GoCrudChallange_Bisrat/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

    // Generate a unique ID for the person
    person.ID = uuid.New().String()

    // Validate required fields
    if person.Name == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
        return
    }
    if person.Age <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Age must be a positive number"})
        return
    }
    if person.Hobbies == nil {
        person.Hobbies = []string{}
    }

    // Create the person
    if err := services.CreatePerson(&person); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, person)
}

func UpdatePerson(c *gin.Context) {
    id := c.Param("id")
    var updatedPerson data.Person
    if err := c.ShouldBindJSON(&updatedPerson); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    
    updatedPerson.ID = id // ID cannot be updated, it will be ignored 

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