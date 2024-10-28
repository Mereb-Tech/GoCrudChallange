package main

import (
	controller "mereb/Controller"
	repository "mereb/Repository"
	router "mereb/Router"
	usecase "mereb/UseCase"

	"github.com/gin-gonic/gin"
)



func main() {
   personRepo := repository.NewPersonRepository()
   personUsecase := usecase.NewPersonUsecase(personRepo)
   personController := controller.NewPersonController(personUsecase)
   
   personRouter := router.NewPersonRouter(personController)
   router := gin.Default()

   personRouter.SetUpRouter(router)
   router.Run(":8080")

}
