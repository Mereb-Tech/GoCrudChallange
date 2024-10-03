package routers

import (
	"GoCrudChallange_Bisrat/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	 // Routes
    r.GET("/person", controllers.GetAllPersons)
    r.GET("/person/:id", controllers.GetPerson)
    r.POST("/person", controllers.CreatePerson)
    r.PUT("/person/:id", controllers.UpdatePerson)
    r.DELETE("/person/:id", controllers.DeletePerson)
	return r
}