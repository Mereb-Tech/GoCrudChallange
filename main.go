package main

import (
	controller "GoCrudChallenge/Controller"
	repository "GoCrudChallenge/Repository"
	routers "GoCrudChallenge/Routers"
	usecases "GoCrudChallenge/UseCases"
	"GoCrudChallenge/config"

	"github.com/gin-gonic/gin"
)



func main(){	
	personRepository := repository.NewPersonRepository()
	personUsecase := usecases.NewPersonUseCase(personRepository)
	personController := controller.NewPersonController(personUsecase)

	router := gin.Default()
	routers.InitializeRouter(router, personController)

	env := config.NewConfig()
	router.Run(env.BASE_URL)

}