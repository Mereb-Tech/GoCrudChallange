package main

import (
	"net/http"

	routers "github.com/poseidon2022/GoCrudChallange/delivery/router"
	domain "github.com/poseidon2022/GoCrudChallange/domain"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	corsConfig := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}

	//I don't get why this is necessary since the browser any test environment
	//handles this by default
	router.NoRoute(func(c *gin.Context) {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"error": "Resource not found",
		})
	})

	router.Use(cors.New(corsConfig))
	routers.PersonRoutes(router, &domain.InMemory)
	router.Run("localhost:3000")
}
