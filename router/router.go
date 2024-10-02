package router

import (
	"time"

	"github.com/abe16s/GoCrudChallange/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(controller controllers.Controller) *gin.Engine {
    r := gin.Default()

    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

	r.POST("/person", controller.CreatePerson)
	r.GET("/person", controller.GetAllPersons)
	r.GET("/person/:id", controller.GetPersonById)
	r.PUT("/person/:id", controller.UpdatePerson)
	r.DELETE("/person/:id", controller.DeletePerson)

    return r
}
