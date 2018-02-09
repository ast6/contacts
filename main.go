package main

import (
	"log"
	"net/http"

	"fmt"

	"github.com/ast6/contacts/Resources"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	Resources.Migration() //create table

	router := mux.NewRouter()
	router.HandleFunc("/add", Resources.Add).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
	fmt.Println(router)
}
