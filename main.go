package main

import(
	"context"
	"encoding/json"
	"log"
	"net/http"

	"helper"
	"models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

var collection = helper.ConnectDB()

func getUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var user models.user
	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&use)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/api/users", createUser).Methods("POST")
	r.HandleFunc("/contacts", newContact).Methods("POST")
	r.HandleFunc("/contacts", newContact).Methods("POST")

	config := helper.GetConfiguration()
	log.Fatal(http.ListenAndServe(config.Port, r))

}