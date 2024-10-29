package Domain

import "github.com/gin-gonic/gin"

type Person struct {
	ID      string   `json:"uid" validate:"required"`
	Name    string   `json:"name" validate:"required,min=2,max=100"`
	Age     int      `json:"age" validate:"required,gte=1,lte=150"`
	Hobbies []string `json:"hobbies" validate:"dive,min=1,max=50"` 
}

type PersonRepository interface {
	GetAllPersons() ([]Person, error)
	GetPersonByID(id string) (Person, error)
	CreatePerson(person Person) (Person, error)
	UpdatePerson(id string, person Person) (Person, error)
	DeletePerson(id string) error
}

type PersonUsecase interface {
	GetAllPersons() ([]Person, error)
	GetPersonByID(id string) (Person, error)
	CreatePerson(person Person) (Person, error)
	UpdatePerson(id string, person Person) (Person, error)
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
