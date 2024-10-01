package main

import (
	routers "mereb_go/delivery/router"
	domain "mereb_go/domain"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routers.PersonRoutes(router, &domain.InMemory)
	router.Run("localhost:3000")
}