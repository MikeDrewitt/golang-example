package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	"golang-example/middleware"
	"golang-example/routes"
)

var port = os.Getenv("PORT")

func start() {
	if port == "" {
		port = "8000"
	}
	router := chi.NewRouter()

	// Generic all routes middleware
	router.Use(middleware.Generic)

	router.Mount("/users", routes.User())
	router.Mount("/status", routes.Status())

	fmt.Printf("Server started and listening at localhost:%s\n", port)

	http.ListenAndServe(":"+port, router)
}
