package main

import(
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/faygun/go-rest-api/helper"
	"github.com/faygun/go-rest-api/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

var collection = helper.ConnectDB()

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/api/users", createUser).Methods("POST")
	r.HandleFunc("/contacts", newContact).Methods("POST")
	//r.HandleFunc("/contacts", newContact).Methods("POST")

	config := helper.GetConfiguration()
	log.Fatal(http.ListenAndServe(config.Port, r))

}