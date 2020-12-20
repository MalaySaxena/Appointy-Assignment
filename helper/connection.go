package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func ConnectDB() *mongo.Client {

	clientOptions := options.Client().ApplyURI("mongodb+srv://msdy:LrjQf9zr5UKZivS@cluster0.gn6vk.mongodb.net/appointy_assignment?retryWrites=true&w=majority")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}


type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}


func GetError(err error, w http.ResponseWriter) {

	log.Fatal(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}
