package main

import(
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"appointy/helper"
	"appointy/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

var client = helper.ConnectDB()

func getUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var user models.User
	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"id": id}
	err := client.Database("appointy_assignment").Collection("user").FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	_ = json.NewDecoder(r.Body).Decode(&user)

	var collection = client.Database("appointy_assignment").Collection("user")
	result, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func getContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var contacts []models.Contact

	var params = mux.Vars(r)

	id := params["user"]
	//timestamp in milliseconds
	timestamp := params["timestamp"]

	id,_ = strconv.ParseInt(id, 10, 64)
	timestamp,_ = strconv.ParseInt(timestamp,10,64)

	var collecti,on = client.Database("appointy_assignment").Collection("contact")

	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var contact models.Contact
		
		err := cur.Decode(&contact) 
		if err != nil {
			log.Fatal(err)
		}
		
		if contact.userIdOne == id || contact.userIdTwo == id {
			var currT, _ = strconv.parse(contact.meetTimestamp,10,64)
			if timestamp-currT <= 1382400000 {
				contacts = append(contacts,contact)
			} 
		} 
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(contacts) 
}

func addContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var contact models.Contact

	_ = json.NewDecoder(r.Body).Decode(&contact)

	var collection = client.Database("appointy_assignment").Collection("contact")
	result, err := collection.InsertOne(context.TODO(), contact)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}


func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/api/users", createUser).Methods("POST")
	r.HandleFunc("/contacts", newContact).Methods("POST")
	r.HandleFunc("/contact/", newContact).Methods("GET")

}