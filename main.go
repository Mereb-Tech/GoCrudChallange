package main

import (
	"GoCrudChallange_Bisrat/routers"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load env variables 
	err:= godotenv.Load()
	if err!=nil{
		 log.Println("Error loading .env file, using default settings")
    
	}
	 port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
	router := routers.SetUpRouter()
	//  CORS middleware
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://example.com", "http://anotherdomain.com"}, // front end domain will be added here 
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))
	
    // Define a custom 404 handler
    router.NoRoute(func(c *gin.Context) {
        c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
    })
	 router.Run(":" + port)


}

