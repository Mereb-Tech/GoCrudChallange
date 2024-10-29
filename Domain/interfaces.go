package Domain

import (
	Models "mereb/Domain/Models"

	"github.com/gin-gonic/gin"
)

type PersonRepository interface {
	GetAllPersons() ([]Models.Person, error)
	GetPersonByID(id string) (Models.Person, error)
	CreatePerson(person Models.Person) (Models.Person, error)
	UpdatePerson(id string, person Models.Person) (Models.Person, error)
	DeletePerson(id string) error
}

type PersonUsecase interface {
	GetAllPersons() ([]Models.Person, error)
	GetPersonByID(id string) (Models.Person, error)
	CreatePerson(person Models.Person) (Models.Person, error)
	UpdatePerson(id string, person Models.Person) (Models.Person, error)
	DeletePerson(id string) error
}
type PersonController interface {
	CreatePerson(ctx *gin.Context)
	GetAllPersons(ctx *gin.Context)
	GetPerson(ctx *gin.Context)
	UpdatePerson(ctx *gin.Context)
	DeletePerson(ctx *gin.Context)
	RouteDoesNotExist(ctx *gin.Context)
}
