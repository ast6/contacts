package main

import (
	"log"
	"net/http"

	"github.com/ast6/contacts/Resources"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	//Resources.Migration() //create table

	router := mux.NewRouter()
	//router.HandleFunc("/add", Resources.Add).Methods("POST")
	router.HandleFunc("/update", Resources.Update).Methods("POST")
	router.HandleFunc("/delete", Resources.Delete).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))

}
