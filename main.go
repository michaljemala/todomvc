package main

import (
	"code.google.com/p/gorilla/mux"
	"net/http"
	"log"
)

func main() {
	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(".")))
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
