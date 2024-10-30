package Domain

import (
	Models "mereb/Domain/Models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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

type InfraStructure interface {
	ValidateStruct(s interface{}) error
	UUID() string
}

type PersonInfrastructure struct {
}

func NewPersonInfrastructure() InfraStructure {
	return &PersonInfrastructure{}
}
func (pv *PersonInfrastructure) ValidateStruct(s interface{}) error {
	return validator.New().Struct(s)
}

func (pv *PersonInfrastructure) UUID() string {
	return uuid.New().String()
}
