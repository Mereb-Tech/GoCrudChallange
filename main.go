package main

import (
	"fmt"
	config "mereb/Config"
	controller "mereb/Controller"
	"mereb/Domain"
	repository "mereb/Repository"
	router "mereb/Router"
	usecase "mereb/UseCase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	personRepo := repository.NewPersonRepository()
	infra := Domain.NewPersonInfrastructure()
	personUsecase := usecase.NewPersonUsecase(personRepo, infra)

	personController := controller.NewPersonController(personUsecase)

	personRouter := router.NewPersonRouter(personController)
	router := gin.Default()

	router.Use(cors.New(config.CorsConfig))

	personRouter.SetUpRouter(router)
	router.Run(":8080")
	fmt.Print("Server running on port 8008")

}
