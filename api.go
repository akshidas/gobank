package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
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
	router.HandleFunc("/accounts", makeHTTPHandleFunc(handlerAccountFunc))
	router.HandleFunc("/accounts/{id}", makeHTTPHandleFunc(handlerAccountFunc))
	http.ListenAndServe(s.port, router)
}

// Handler for root route
func getRoot(w http.ResponseWriter, r *http.Request) error {
	fmt.Printf("got / request\n")
	return writeJson(w, http.StatusOK, "Server UP and running")
}
