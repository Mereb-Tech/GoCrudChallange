package controller

import (
	"mereb/Domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PersonController struct {
	usecase Domain.PersonUsecase
}

func NewPersonController(usecase Domain.PersonUsecase) *PersonController {
	return &PersonController{usecase}
}

func (controller *PersonController) GetAllPersons(ctx *gin.Context) {
	persons, err := controller.usecase.GetAllPersons()
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, persons)
}

func (controler *PersonController) GetPerson(ctx *gin.Context) {
	id := ctx.Param("id")
	person, err := controler.usecase.GetPersonByID(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, person)
}

func (controller *PersonController) CreatePerson(ctx *gin.Context) {
	var person Domain.Person
	if err := ctx.BindJSON(&person); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	person, err := controller.usecase.CreatePerson(person)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusCreated, person)
}

func (controller *PersonController) UpdatePerson(ctx *gin.Context) {
	id := ctx.Param("id")
	var person Domain.Person
	if err := ctx.BindJSON(&person); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	person, err := controller.usecase.UpdatePerson(id, person)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, person)
}

func (controller *PersonController) DeletePerson(ctx *gin.Context) {
	id := ctx.Param("id")
	err := controller.usecase.DeletePerson(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "person deleted successfully"})

}
