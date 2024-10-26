package middlewares

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
    c.Next() 
	
    if len(c.Errors) > 0 {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
    }
}
