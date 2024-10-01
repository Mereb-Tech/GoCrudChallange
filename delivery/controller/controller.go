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

func (pc *PersonController) UpdatePerson() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}


func (pc *PersonController) DeletePerson() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
