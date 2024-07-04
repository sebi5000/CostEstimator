package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	router.Get("/", indexHandler)

	router.Post("/calculate_price", calculatePriceHandler)
	router.Post("/clear_price", clearPriceHandler)

	http.ListenAndServe(":8000", router)
}
