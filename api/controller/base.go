package controller

import (
	errapi "github.com/beka-birhanu/GoCrudChallange/api/error"
	"github.com/gin-gonic/gin"
)

// BaseController is a base struct for all HTTP request handlers that provides basic HTTP functionalities.
type BaseController struct{}

// RespondError handles errors by writing an appropriate response to the Gin context.
// The primary usage is to hide specific messages from the client.
func (h *BaseController) RespondError(c *gin.Context, err errapi.Error) {
	if err.StatusCode() == errapi.ServerError {
		err = errapi.NewServerError("something went wrong")
	}
	c.JSON(err.StatusCode(), gin.H{"error": err.Error()})
}

// Respond writes a JSON response to the Gin context.
func (h *BaseController) Respond(c *gin.Context, status int, v interface{}) {
	if v == nil {
		c.Status(status)
	} else {
		c.JSON(status, v)
	}
}

// RespondWithLocation writes a response with a Location header to the Gin context.
func (h *BaseController) RespondWithLocation(c *gin.Context, status int, v interface{}, resourceLocation string) {
	c.Header("Location", resourceLocation)
	h.Respond(c, status, v)
}
