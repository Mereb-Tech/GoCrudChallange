package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Mahider-T/GoCrudChallange/internal/adapters/db"
	"github.com/Mahider-T/GoCrudChallange/internal/adapters/handlers"
	"github.com/Mahider-T/GoCrudChallange/internal/application/core/api"
	"github.com/Mahider-T/GoCrudChallange/internal/application/core/domain"
	"github.com/Mahider-T/GoCrudChallange/internal/ports/apiport"
)

type Application struct {
	// userRepo    dbport.UserDbPort
	userService apiport.UserApiPort
}

func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	mux := http.NewServeMux()

	var storage = make(map[string]domain.User)

	app := &Application{
		userService: api.NewUserApi(db.NewUserRepo(storage)),
	}

	userHandler := handlers.NewUserHandler(app.userService)

	mux.HandleFunc("POST /users", userHandler.UserCreateHandler)
	mux.HandleFunc("GET /users", userHandler.UserGetAllHandler)
	mux.HandleFunc("GET /users/{id}", userHandler.UserGetByIdHandler)
	mux.HandleFunc("PUT /users/{id}", userHandler.UserUpdateHandler)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Println("Listening on port:", port)
	http.ListenAndServe(":"+port, enableCors(mux))

}
