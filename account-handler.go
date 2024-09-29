package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func handlerAccountFunc(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return getAccount(w, r)
	}

	if r.Method == "POST" {
		return AddAccount(w, r)
	}

	if r.Method == "PUT" {
		return UpdateAccount(w, r)
	}
	return writeJson(w, http.StatusMethodNotAllowed, "method not allowed")
}

// getAccount will handle get request for fetch all and for fetch by id
func getAccount(w http.ResponseWriter, r *http.Request) error {
	accountId, ok := mux.Vars(r)["id"]
	if !ok {
		return writeJson(w, http.StatusOK, DB.getAll())
	}

	return GetAccountById(accountId, w, r)
}

func GetAccountById(accountId string, w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(accountId)
	if err != nil {
		return writeJson(w, http.StatusInternalServerError, err)
	}

	account, _ := DB.getById(int(id))
	if account == nil {
		return writeJson(w, http.StatusNotFound, "Not Found")
	}

	return writeJson(w, http.StatusOK, account)
}

func AddAccount(w http.ResponseWriter, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	t := Account{}
	decoder.Decode(&t)
	DB.Add(&t)
	return writeJson(w, http.StatusCreated, "create route")
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) error {
	accountId, ok := mux.Vars(r)["id"]
	if !ok {
		return writeJson(w, http.StatusNotFound, "Not Found")
	}

	id, err := strconv.Atoi(accountId)
	if err != nil {
		return writeJson(w, http.StatusInternalServerError, err)
	}

	decoder := json.NewDecoder(r.Body)
	t := Account{}
	t.ID = id
	decoder.Decode(&t)
	DB.Update(id, t)
	return writeJson(w, http.StatusCreated, "update route")
}
