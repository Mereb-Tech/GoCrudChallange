package routers

import (
	interfaces "GoCrudChallenge/Domain/Interfaces"
	"GoCrudChallenge/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)




func InitializeRouter(router *gin.Engine, personController interfaces.PersonControllerInterface){

    router.Use(middlewares.ErrorHandler)
    router.Use(middlewares.CORSMiddleware())

	v1 := router.Group("/api/v1")

	personRoutes := v1.Group("/person")
	{
		personRoutes.POST("/", personController.CreatePerson)
		personRoutes.GET("/", personController.GetAllPersons)
		personRoutes.GET("/:id", personController.GetPersonByID)
		personRoutes.PUT("/:id", personController.UpdatePerson)
		personRoutes.DELETE("/:id", personController.DeletePerson)
	}

	router.NoRoute(func(c *gin.Context) {
        c.JSON(http.StatusNotFound, gin.H{"error": "resource not found"})
    })
}