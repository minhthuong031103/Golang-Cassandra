package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/create", CreateUser).Methods("POST")         // http://localhost:3000/create
	router.HandleFunc("/user/all", GetAllUsers).Methods("GET")       // http://localhost:3000/getstudents
	router.HandleFunc("/count", CountAllUsers).Methods("GET")        // http://localhost:3000/count
	router.HandleFunc("/getone/{userId}", GetOneUser).Methods("GET") // http://localhost:3000/getone/1

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(headers, methods, origins)(router)))

}
