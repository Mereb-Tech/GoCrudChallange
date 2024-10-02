package router

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/poseidon2022/GoCrudChallange/delivery/controller"
	domain "github.com/poseidon2022/GoCrudChallange/domain"
	"github.com/poseidon2022/GoCrudChallange/repository"
	"github.com/poseidon2022/GoCrudChallange/usecases"
)

func PersonRoutes(internalRouter *gin.Engine, persons *[]domain.Person) {
	pr := repository.NewPersonRepository(persons)
	pc := &controllers.PersonController{
		PersonUseCase: usecases.NewPersonUseCase(pr),
	}

	internalRouter.GET("/api/person", pc.GetAllPersons())
	internalRouter.GET("/api/person/:person_id", pc.GetPersonById())
	internalRouter.POST("/api/person", pc.Register())
	internalRouter.PUT("/api/person/:person_id", pc.UpdatePerson())
	internalRouter.DELETE("/api/person/:person_id", pc.DeletePerson())
}
