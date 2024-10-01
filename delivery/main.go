package main

import (
	routers "mereb_go/delivery/router"
	domain "mereb_go/domain"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	router := gin.Default()

	corsConfig := cors.Config{
        AllowAllOrigins: true,
        AllowMethods:    []string{"GET", "POST", "PUT", "DELETE"}, 
        AllowHeaders:   []string{"Origin", "Content-Type", "Accept"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true, 
    }

    router.Use(cors.New(corsConfig))
	routers.PersonRoutes(router, &domain.InMemory)
	router.Run("localhost:3000")
}