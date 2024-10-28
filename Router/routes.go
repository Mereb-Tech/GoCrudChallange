package router

import (
	"mereb/Domain"

	"github.com/gin-gonic/gin"
)

type Router struct {
	controller Domain.PersonController
}

func NewPersonRouter(contoller Domain.PersonController) *Router {
	return &Router{
		controller: contoller,
	}
}

func (route *Router) SetUpRouter(router *gin.Engine) {
	personRouter := router.Group("/person")
	{
		personRouter.POST("", route.controller.CreatePerson)
		personRouter.GET("", route.controller.GetAllPersons)
		personRouter.GET("/:id", route.controller.GetPerson)
		personRouter.PUT("/:id", route.controller.UpdatePerson)
		personRouter.DELETE("/:id", route.controller.DeletePerson)
	}

}
