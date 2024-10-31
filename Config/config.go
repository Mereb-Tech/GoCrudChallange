package config

import "github.com/gin-contrib/cors"

var ErrorMessages = map[string]string{
	"Person.ID.required":   "User ID is required.",
	"Person.Name.required": "Name is required.",
	"Person.Name.min":      "Name must be at least 2 characters.",
	"Person.Name.max":      "Name can be a maximum of 100 characters.",
	"Person.Age.required":  "Age is required.",
	"Person.Age.gte":       "Age must be at least 1 year.",
	"Person.Age.lte":       "Age must be less than or equal to 150 years.",
	"Person.Hobbies.min":   "Each hobby must be a non-empty string.",
	"Person.Hobbies.max":   "Each hobby can be a maximum of 50 characters.",
}

var CorsConfig = cors.Config{
	AllowAllOrigins: true, // Allows all orgins
	AllowHeaders:    []string{"Origin", "Content-Type", "Accept"},
	AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	MaxAge:          12 * 3600,
}
