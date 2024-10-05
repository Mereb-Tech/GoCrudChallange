package main

import (
	"fmt"
	"log"

	"github.com/beka-birhanu/GoCrudChallange/api/controller"
	"github.com/beka-birhanu/GoCrudChallange/api/router"
	personcmd "github.com/beka-birhanu/GoCrudChallange/app/person/commands"
	personqry "github.com/beka-birhanu/GoCrudChallange/app/person/query"
	"github.com/beka-birhanu/GoCrudChallange/config"
	"github.com/beka-birhanu/GoCrudChallange/infrastructure/repo"
)

func main() {
	cfg := config.Envs

	// Initialize repository and handlers
	personRepo := repo.NewPersonRepo()
	createPersonHandler := personcmd.NewCreatePersonHandler(personRepo)
	updatePersonHandler := personcmd.NewUpdatePersonHandler(personRepo)
	deletePersonHandler := personcmd.NewDeletePersonHandler(personRepo)
	getPersonHandler := personqry.NewGetPersonHandler(personRepo)
	getAllPersonsHandler := personqry.NewGetAllPersonHandler(personRepo)

	// Initialize controller
	personController := controller.NewPersonController(controller.NewPersonControllerParams{
		CreateHandler: createPersonHandler,
		UpdateHandler: updatePersonHandler,
		DeleteHandler: deletePersonHandler,
		GetHandler:    getPersonHandler,
		GetAllHandler: getAllPersonsHandler,
	})

	// Initialize router
	controllers := []controller.IController{personController}
	r := router.NewRouter(router.Config{
		Addr:        fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.ServerPort),
		Controllers: controllers,
	})

	// Run the server
	if err := r.Run(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
