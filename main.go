package main

import (
	"go-templ-app/handlers"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	router.Get("/", handlers.HomeHandler)
	router.Post("/greet", handlers.HandleGreet)

	// Melayani file statis
	fs := http.FileServer(http.Dir("./static"))
	router.Handle("/static/*", http.StripPrefix("/static/", fs))

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
