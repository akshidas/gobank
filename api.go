package main

import (
	"fmt"
	"github.com/gorilla/mux"
	account "gobank/accounts"
	"net/http"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type ApiError struct {
	Error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			account.WriteJson(w, http.StatusBadRequest, &ApiError{Error: err.Error()})
		}
	}
}

type ApiServer struct {
	port string
}

func NewApiServer(listentPort string) *ApiServer {
	return &ApiServer{
		port: listentPort,
	}
}

func (s *ApiServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/", makeHTTPHandleFunc(getRoot))
	router.HandleFunc("/accounts", makeHTTPHandleFunc(account.HandlerAccountFunc))
	router.HandleFunc("/accounts/{id}", makeHTTPHandleFunc(account.HandlerAccountFunc))
	http.ListenAndServe(s.port, router)
}

// Handler for root route
func getRoot(w http.ResponseWriter, r *http.Request) error {
	fmt.Printf("got / request\n")
	return account.WriteJson(w, http.StatusOK, "Server UP and running")
}
