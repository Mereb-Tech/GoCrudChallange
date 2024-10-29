package main

import (
	"fmt"
	config "mereb/Config"
	controller "mereb/Controller"
	repository "mereb/Repository"
	router "mereb/Router"
	usecase "mereb/UseCase"

	"github.com/gin-contrib/cors"	
	"github.com/gin-gonic/gin"
)

func main() {
	personRepo := repository.NewPersonRepository()
	personUsecase := usecase.NewPersonUsecase(personRepo)
	personController := controller.NewPersonController(personUsecase)

	personRouter := router.NewPersonRouter(personController)
	router := gin.Default()

	router.Use(cors.New(config.CorsConfig))

	personRouter.SetUpRouter(router)
	router.Run(":8080")
	fmt.Print("Server running on port 8008")

}
