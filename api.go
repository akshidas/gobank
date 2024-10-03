package main

import (
	"fmt"
	"gobank/accounts"
	"gobank/helpers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type ApiError struct {
	Error string
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
	router.HandleFunc("/", helpers.MakeHTTPHandleFunc(getRoot))
	router.HandleFunc("/accounts", helpers.MakeHTTPHandleFunc(accounts.HandlerAccountFunc))
	router.HandleFunc("/accounts/{id}", helpers.MakeHTTPHandleFunc(accounts.HandlerAccountFunc))
	log.Printf("Starting the server on port%s", s.port)
	err := http.ListenAndServe(s.port, router)
	if err != nil {
		log.Fatalf("Failed to start the server due to: %s", err)
	}
}

// Handler for root route
func getRoot(w http.ResponseWriter, r *http.Request) error {
	fmt.Printf("got / request\n")
	return helpers.WriteJson(w, http.StatusOK, "Server UP and running")
}
