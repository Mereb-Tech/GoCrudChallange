// cmd/main.go
package main

import (
	"log"
	"mereb_assessment/controllers"
	"mereb_assessment/repositories"
	"mereb_assessment/services"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := repositories.NewPersonRepository()
	service := services.NewPersonService(repo)
	handler := handlers.NewPersonHandler(service)

	router := gin.Default()

	handler.RegisterRoutes(router)

	log.Println("Server is running on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
