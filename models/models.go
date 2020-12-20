package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Create Struct
type User struct {
	id     primitive.ObjectID `json:"id" bson:"id"`
	name   string  `json:"name" bson:"name"`
	dateOfBirth  string `json:"date_of_birth" bson:"date_of_birth"`
	phoneNumber string `json:"phone_number" bson:"phone_number"`
	emailAddress string `json:"email_address" bson:"email_address"`
	createTime string `json: "create_timestamp" bson:"create_timestamp"`
}

type Contact struct {
	userIdOne int `json:"user1_id" bson:"user1_id"`
	userIdTwo int `json:"user2_id" bson:"user2_id"`
	meetTimestamp string `json:"meet_timestamp" bson:"meet_timestamp"`
}