package routes

import (
	"golang-example/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Status() http.Handler {
	router := chi.NewRouter()

	router.Get("/", statusGet)

	return router
}

func statusGet(w http.ResponseWriter, r *http.Request) {
	utils.Ok(w, "Hello, the API is working")
}
