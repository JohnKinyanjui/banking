package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	mux := mux.NewRouter()

	mux.HandleFunc("/greet", greet).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
