package main

import (
	"github.com/abe16s/GoCrudChallange/controllers"
	"github.com/abe16s/GoCrudChallange/repositories"
	"github.com/abe16s/GoCrudChallange/router"
	"github.com/abe16s/GoCrudChallange/usecases"
)

func main() {
	repo := repositories.NewRepository()
	service := usecases.NewService(repo)
	controller := controllers.NewController(service)

	r := router.NewRouter(*controller)
	r.Run("localhost:8080")
}
