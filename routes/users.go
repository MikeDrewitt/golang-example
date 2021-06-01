package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"golang-example/middleware"
	"golang-example/models"

	"github.com/go-chi/chi/v5"
)

func User() http.Handler {
	router := chi.NewRouter()

	router.Get("/", userList)
	router.Get("/{id}", userGet)
	router.Post("/", userPost)
	router.Put("/{id}", userPut)
	router.Delete("/{id}", userDelete)

	return router
}

func userGet(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	intId, err := strconv.Atoi(id)

	if err != nil {
		middleware.BadRequest(w, "Bad Id")
		return
	}

	user, dbErr := models.UserGet(intId)

	if dbErr != nil {
		middleware.NotFound(w, dbErr.Error())
		return
	}

	middleware.Ok(w, user)
}

func userList(w http.ResponseWriter, r *http.Request) {
	users, dbErr := models.UserList()

	if dbErr != nil {
		middleware.BadRequest(w, dbErr.Error())
		return
	}

	middleware.Ok(w, users)
}

func userPost(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var newUser models.User
	json.Unmarshal(reqBody, &newUser)

	validationErr := models.UserValidation(newUser)

	if validationErr != nil {
		middleware.BadRequest(w, validationErr.Error())
		return
	}

	newUser, dbErr := models.UserCreate(newUser)

	if dbErr != nil {
		middleware.BadRequest(w, dbErr.Error())
		return
	}

	middleware.Ok(w, newUser)
}

func userPut(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	intId, err := strconv.Atoi(id)

	if err != nil {
		middleware.NotFound(w, "Bad Id")
		return
	}

	reqBody, _ := ioutil.ReadAll(r.Body)

	var updatedUser models.User
	json.Unmarshal(reqBody, &updatedUser)

	if intId != updatedUser.Id {
		middleware.BadRequest(w, "Id in URL must match Id in body")
		return
	}

	validationErr := models.UserValidation(updatedUser)

	if validationErr != nil {
		middleware.BadRequest(w, validationErr.Error())
		return
	}

	updatedUser, dbErr := models.UserUpdate(updatedUser)

	if dbErr != nil {
		middleware.BadRequest(w, dbErr.Error())
		return
	}

	middleware.Ok(w, updatedUser)
}

func userDelete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	intId, err := strconv.Atoi(id)

	if err != nil {
		middleware.BadRequest(w, "Bad Id")
		return
	}

	dbErr := models.UserDelete(intId)

	if dbErr != nil {
		middleware.NotFound(w, dbErr.Error())
		return
	}

	middleware.NoContent(w)
}
