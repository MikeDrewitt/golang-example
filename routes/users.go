package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"golang-example/models"
	"golang-example/utils"

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
		utils.BadRequest(w, "Bad Id")
		return
	}

	user, err := models.UserGet(intId)
	if err != nil {
		utils.NotFound(w, err.Error())
		return
	}

	utils.Ok(w, user)
}

func userList(w http.ResponseWriter, r *http.Request) {
	users, err := models.UserList()
	if err != nil {
		utils.BadRequest(w, err.Error())
		return
	}

	utils.Ok(w, users)
}

func userPost(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var newUser models.User
	json.Unmarshal(reqBody, &newUser)

	err := models.UserValidation(newUser)
	if err != nil {
		utils.BadRequest(w, err.Error())
		return
	}

	newUser, err = models.UserCreate(newUser)
	if err != nil {
		utils.BadRequest(w, err.Error())
		return
	}

	utils.Ok(w, newUser)
}

func userPut(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		utils.NotFound(w, "Bad Id")
		return
	}

	reqBody, _ := ioutil.ReadAll(r.Body)

	var updatedUser models.User
	json.Unmarshal(reqBody, &updatedUser)

	if intId != updatedUser.Id {
		utils.BadRequest(w, "Id in URL must match Id in body")
		return
	}

	err = models.UserValidation(updatedUser)
	if err != nil {
		utils.BadRequest(w, err.Error())
		return
	}

	updatedUser, err = models.UserUpdate(updatedUser)
	if err != nil {
		utils.BadRequest(w, err.Error())
		return
	}

	utils.Ok(w, updatedUser)
}

func userDelete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		utils.BadRequest(w, "Bad Id")
		return
	}

	err = models.UserDelete(intId)
	if err != nil {
		utils.NotFound(w, err.Error())
		return
	}

	utils.NoContent(w)
}
