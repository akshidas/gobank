package accounts

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gobank/helpers"
	"math/rand"
	"net/http"
	"strconv"
)

func HandlerAccountFunc(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return getAccount(w, r)
	}

	if r.Method == "POST" {
		return addAccount(w, r)
	}

	if r.Method == "PUT" {
		return updateAccount(w, r)
	}
	return helpers.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
}

// getAccount will handle get request for fetch all and for fetch by id
func getAccount(w http.ResponseWriter, r *http.Request) error {
	accountId, ok := mux.Vars(r)["id"]
	if !ok {
		return helpers.WriteJson(w, http.StatusOK, getAll())
	}
	return getAccountById(accountId, w)
}

func getAccountById(accountId string, w http.ResponseWriter) error {
	id, err := strconv.Atoi(accountId)
	if err != nil {
		return helpers.WriteJson(w, http.StatusInternalServerError, err)
	}

	account, index := getById(int(id))
	if index == -1 {
		return helpers.WriteJson(w, http.StatusNotFound, "Not Found")
	}

	return helpers.WriteJson(w, http.StatusOK, account)
}

func addAccount(w http.ResponseWriter, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	t := Account{}
	decoder.Decode(&t)
	t.ID = rand.Intn(1000)
	Add(t)
	return helpers.WriteJson(w, http.StatusCreated, "create route")
}

func updateAccount(w http.ResponseWriter, r *http.Request) error {
	accountId, ok := mux.Vars(r)["id"]
	if !ok {
		return helpers.WriteJson(w, http.StatusNotFound, "Not Found")
	}

	id, err := strconv.Atoi(accountId)
	if err != nil {
		return helpers.WriteJson(w, http.StatusInternalServerError, err)
	}

	decoder := json.NewDecoder(r.Body)
	t := Account{}
	t.ID = id
	decoder.Decode(&t)
	Update(id, t)
	return helpers.WriteJson(w, http.StatusCreated, "update route")
}
