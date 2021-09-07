package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id              primitive.ObjectID `json:"_id" bson:"_id"`
	Username        string             `json:"username" bson:"username"`
	Email           string             `json:"email" bson:"email"`
	Roles           []Role             `json:"roles" bson:"roles"`
	Profile_picture string             `json:"profile_picture" bson:"profile_picture"`
	Bio             string             `json:"bio" bson:"bio"`
	Display_name    string             `json:"display_name" bson:"display_name"`
	Position        string             `json:"position" bson:"position"`
}
