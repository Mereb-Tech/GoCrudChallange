package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Mahider-T/GoCrudChallange/internal/adapters/db"
	"github.com/Mahider-T/GoCrudChallange/internal/adapters/handlers"
	"github.com/Mahider-T/GoCrudChallange/internal/application/core/api"
	"github.com/Mahider-T/GoCrudChallange/internal/application/core/domain"
	"github.com/Mahider-T/GoCrudChallange/internal/ports/apiport"
)

type Application struct {
	personService apiport.PersonApiPort
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "404 Not Found - Resource does not exist", http.StatusNotFound)
}

func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
func main() {

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	var storage = make(map[string]domain.Person)

	app := &Application{
		personService: api.NewPersonApi(db.NewPersonRepo(storage)),
	}

	personHandler := handlers.NewPersonHandler(app.personService)

	mux.HandleFunc("POST /person", personHandler.PersonCreateHandler)
	mux.HandleFunc("GET /person", personHandler.PersonGetAllHandler)
	mux.HandleFunc("GET /person/{personId}", personHandler.PersonGetByIdHandler)
	mux.HandleFunc("PUT /person/{personId}", personHandler.PersonUpdateHandler)
	mux.HandleFunc("DELETE /person/{personId}", personHandler.PersonDeleteHandler)

	mux.HandleFunc("/", NotFoundHandler)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	infoLog.Printf("Starting server on %s", port)
	addr := fmt.Sprintf(":%s", port)

	ser := &http.Server{

		Addr:         addr,
		Handler:      enableCors(mux),
		ErrorLog:     errorLog,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	err := ser.ListenAndServe()
	errorLog.Fatal(err)
}
