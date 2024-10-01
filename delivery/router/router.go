package router

import (
	"github.com/gin-gonic/gin"
	controllers "mereb_go/delivery/controller"
	domain "mereb_go/domain"
	"mereb_go/usecases"
	"mereb_go/repository"
)

func PersonRoutes(internalRouter *gin.Engine, persons *[]domain.Person) {
	pr := repository.NewPersonRepository(persons)
	pc := &controllers.PersonController{
		PersonUseCase : usecases.NewPersonUseCase(pr),
	}

	internalRouter.GET("/api/person", pc.GetAllPersons())
	internalRouter.GET("/api/person/:person_id", pc.GetPersonById())
	internalRouter.POST("/api/person", pc.Register())
	internalRouter.PUT("/api/person/:person_id", pc.UpdatePerson())
	internalRouter.DELETE("/api/person/:person_id", pc.DeletePerson())
}