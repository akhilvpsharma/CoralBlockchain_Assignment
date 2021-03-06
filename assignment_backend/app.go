package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func Serve() {

	router := mux.NewRouter()
	
	router.HandleFunc("/save", SaveController).Methods("POST")
	router.HandleFunc("/search/{emailId}", SearchController).Methods("GET")
	router.HandleFunc("/delete/{emailId}", DeleteController).Methods("GET")
	
	fmt.Println("Listening on http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", router))
}
