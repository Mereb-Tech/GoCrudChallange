package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"mereb_go/domain"
)

type PersonController struct {
	PersonUseCase 	domain.PersonUseCase
}

func (pc *PersonController) Register() gin.HandlerFunc {

	var newPerson domain.NewPerson
	return func(c *gin.Context) {
		if err := c.BindJSON(&newPerson); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"success" : false,
				"message" :  "invalid person format",
			})
			return
		}

		err := pc.PersonUseCase.Register(&newPerson)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"success" : false,
				"message" : "internal server error",
			})

			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{
			"success" : true,
			"message" : "person registerd successfully",
		})
	}
}

func (pc *PersonController) GetAllPersons() gin.HandlerFunc {
	return func(c *gin.Context) {
		allPersons, err := pc.PersonUseCase.GetAllPersons()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"success" : false,
				"message" : "internal server error",
			})
			return
		}
		
		c.IndentedJSON(http.StatusOK, gin.H{
			"success" : true,
			"message" : "persons fetched successfully",
			"persons" : allPersons,
		})
	}
}

func (pc *PersonController) GetPersonById() gin.HandlerFunc {
	return func(c *gin.Context) {
		person_id := c.Param("person_id")
		foundPerson, err := pc.PersonUseCase.GetPersonById(person_id)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{
				"success" : false,
				"message" : "person with the specified id not found",
			})
			
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{
			"success" : true,
			"message" : "person found and returned successfully",
			"person" : foundPerson,
		})
	}
}

func (pc *PersonController) UpdatePerson() gin.HandlerFunc {
	return func(c *gin.Context) {
		person_id := c.Param("person_id")
		var updatedInfo domain.NewPerson
		if err := c.BindJSON(&updatedInfo); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"success" : false,
				"message" : "invalid request format",
			})

			return
		}
		updatedPerson, err := pc.PersonUseCase.UpdatePerson(updatedInfo, person_id)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{
				"success" : false,
				"message" : "person with the specified id not found",
			})
			
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{
			"success" : true,
			"message" : "person info updated successfully",
			"updated person" : updatedPerson,
		})
	}
}


func (pc *PersonController) DeletePerson() gin.HandlerFunc {
	return func(c *gin.Context) {
		person_id := c.Param("person_id")
		deletedPerson, err := pc.PersonUseCase.DeletePerson(person_id)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{
				"success" : false,
				"message" : "person with the specified id not found",
			})
			
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{
			"success" : true,
			"message" : "person with the specified id found and deleted",
			"deleted_person" : deletedPerson,
		})
	}
}
