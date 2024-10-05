// Package router provides functionality to set up and run the HTTP server,
// manage routes, and apply middleware based on access levels.
//
// It configures and initializes routes with varying access requirements:
// - Public routes: Accessible without authentication.
// - Protected routes: Require authentication.
// - Privileged routes: Require both authentication and admin privileges.
package router

import (
	"log"
	"time"

	"github.com/beka-birhanu/GoCrudChallange/api/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Router manages the HTTP server and its dependencies,
// including controllers and JWT authentication.
type Router struct {
	addr        string
	controllers []controller.IController
}

// Config holds configuration settings for creating a new Router instance.
type Config struct {
	Addr        string                   // Address to listen on
	Controllers []controller.IController // List of controllers
}

// NewRouter creates a new Router instance with the given configuration.
// It initializes the router with address, base URL, controllers, and JWT service.
func NewRouter(config Config) *Router {
	return &Router{
		addr:        config.Addr,
		controllers: config.Controllers,
	}
}

// Run starts the HTTP server and sets up routes with different access levels.
func (r *Router) Run() error {
	router := gin.Default()

	corsConfig := cors.New(cors.Config{
		AllowAllOrigins:  true, // TODO: change this with white list when domain is known
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	})

	router.Use(corsConfig)

	api := router.Group("/api")
	{
		// Public routes (accessible without authentication)
		publicRoutes := api.Group("/v1")
		{
			for _, c := range r.controllers {
				c.RegisterPublic(publicRoutes)
			}
		}

		// Register Protected and Privileged routes if necessary.
	}

	log.Println("Listening on", r.addr)
	if err := router.Run(r.addr); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
	return nil
}
