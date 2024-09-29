package accounts

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

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
	return WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
}

// getAccount will handle get request for fetch all and for fetch by id
func getAccount(w http.ResponseWriter, r *http.Request) error {
	accountId, ok := mux.Vars(r)["id"]
	if !ok {
		return WriteJson(w, http.StatusOK, getAll())
	}

	return getAccountById(accountId, w, r)
}

func getAccountById(accountId string, w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(accountId)
	if err != nil {
		return WriteJson(w, http.StatusInternalServerError, err)
	}

	account, index := getById(int(id))
	if index == -1 {
		return WriteJson(w, http.StatusNotFound, "Not Found")
	}

	return WriteJson(w, http.StatusOK, account)
}

func addAccount(w http.ResponseWriter, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	t := Account{}
	decoder.Decode(&t)
	t.ID = rand.Intn(1000)
	Add(t)
	return WriteJson(w, http.StatusCreated, "create route")
}

func updateAccount(w http.ResponseWriter, r *http.Request) error {
	accountId, ok := mux.Vars(r)["id"]
	if !ok {
		return WriteJson(w, http.StatusNotFound, "Not Found")
	}

	id, err := strconv.Atoi(accountId)
	if err != nil {
		return WriteJson(w, http.StatusInternalServerError, err)
	}

	decoder := json.NewDecoder(r.Body)
	t := Account{}
	t.ID = id
	decoder.Decode(&t)
	Update(id, t)
	return WriteJson(w, http.StatusCreated, "update route")
}
