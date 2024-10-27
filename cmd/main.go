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
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint"},
	)
)

func init() {
	prometheus.MustRegister(requestCount)
}

type application struct {
	personService apiport.PersonApiPort
}

func main() {

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()
	var storage = make(map[string]domain.Person)

	app := &application{
		personService: api.NewPersonApi(db.NewPersonRepo(storage)),
	}

	personHandler := handlers.NewPersonHandler(app.personService)

	mux.Handle("GET /metrics", promhttp.Handler())
	mux.HandleFunc("POST /person", app.MetricsMiddleware(personHandler.PersonCreateHandler, "person_create"))
	mux.HandleFunc("GET /person", app.MetricsMiddleware(personHandler.PersonGetAllHandler, "person_get_all"))
	mux.HandleFunc("GET /person/{personId}", app.MetricsMiddleware(personHandler.PersonGetByIdHandler, "person_get_by_id"))
	mux.HandleFunc("PUT /person/{personId}", app.MetricsMiddleware(personHandler.PersonUpdateHandler, "person_update"))
	mux.HandleFunc("DELETE /person/{personId}", app.MetricsMiddleware(personHandler.PersonDeleteHandler, "person_delete"))

	mux.HandleFunc("/", personHandler.NotFoundHandler)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	infoLog.Printf("Starting server on %s", port)
	addr := fmt.Sprintf(":%s", port)

	ser := &http.Server{

		Addr:         addr,
		Handler:      app.enableCors(mux),
		ErrorLog:     errorLog,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	err := ser.ListenAndServe()
	errorLog.Fatal(err)
}

func (app *application) MetricsMiddleware(handler http.HandlerFunc, handlerName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestCount.WithLabelValues(r.Method, handlerName).Inc()
		handler(w, r)
	}
}

func (app *application) enableCors(next http.Handler) http.Handler {
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
