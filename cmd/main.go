package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	router.Get("/", indexHandler)
	router.Get("/request-calculation", requestCalculationHandler)

	router.Post("/calculate_price", calculatePriceHandler)
	router.Post("/clear_price", clearPriceHandler)

	router.Post("/calculate_request", calculateRequestHandler)
	router.Post("/clear_request", clearRequestHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}
