package main

import (
	controller "mereb/Controller"
	repository "mereb/Repository"
	router "mereb/Router"
	usecase "mereb/UseCase"

	"github.com/gin-gonic/gin"
   "github.com/gin-contrib/cors"
)



func main() {
   personRepo := repository.NewPersonRepository()
   personUsecase := usecase.NewPersonUsecase(personRepo)
   personController := controller.NewPersonController(personUsecase)
   
   personRouter := router.NewPersonRouter(personController)
   router := gin.Default()

   corsConfig := cors.Config{
		AllowAllOrigins: true,  // Allows all orgins
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		MaxAge:           12 * 3600, 
	}
   router.Use(cors.New(corsConfig))

   personRouter.SetUpRouter(router)
   router.Run(":8080")

}
