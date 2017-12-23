package main

import (
	"StatusCheck/api"
	"StatusCheck/database"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func connectDB(db database.DB) {
	fmt.Printf("Connected: %v\n", db.Connect())

}

func main() {
	test := &database.Redis{}
	test.IP = "localhost"
	test.Port = "6379"
	test.DB = 0
	test.Password = ""
	connectDB(test)

	api := api.API{
		Database: test,
	}

	r := mux.NewRouter()

	r.HandleFunc("/", api.IndexHandler).Methods("GET")
	r.HandleFunc("/Live/{url}", api.Live).Methods("GET")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.ListenAndServe("0.0.0.0:8000", r)

}
