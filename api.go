package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type ApiError struct {
	Error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			writeJson(w, http.StatusBadRequest, &ApiError{Error: err.Error()})
		}
	}
}

func writeJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type ApiServer struct {
	port string
}

func NewApiServer(listentPort string) *ApiServer {
	return &ApiServer{
		port: listentPort,
	}
}

var DB = DataBase{
	accounts: []Account{{ID: 1, FirstName: "Akshay", LastName: "Krishna", Number: 343, Balance: 34}},
}

func (s *ApiServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/", makeHTTPHandleFunc(getRoot))
	router.HandleFunc("/account", makeHTTPHandleFunc(handlerAccountFunc))
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(handlerAccountFunc))
	http.ListenAndServe(s.port, router)
}

// Handler for root route
func getRoot(w http.ResponseWriter, r *http.Request) error {
	fmt.Printf("got / request\n")
	return writeJson(w, http.StatusOK, "Server UP and running")
}

func handlerAccountFunc(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return getAccount(w, r)
	}

	if r.Method == "POST" {
		return AddAccount(w, r)
	}
	return writeJson(w, http.StatusMethodNotAllowed, "method not allowed")
}

// getAccount will handle get request for fetch all and for fetch by id
func getAccount(w http.ResponseWriter, r *http.Request) error {
	accountId, ok := mux.Vars(r)["id"]
	if !ok {
		return writeJson(w, http.StatusOK, DB.getAccounts())
	}

	id, err := strconv.Atoi(accountId)
	if err != nil {
		return writeJson(w, http.StatusInternalServerError, err)
	}

	account := DB.getAccountById(int(id))
	if account == nil {
		return writeJson(w, http.StatusNotFound, "Not Found")
	}

	return writeJson(w, http.StatusOK, account)
}

func AddAccount(w http.ResponseWriter, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	t := Account{}
	decoder.Decode(&t)
	DB.AddAccount(&t)
	return writeJson(w, http.StatusCreated, "create route")
}
