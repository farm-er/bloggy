package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"`
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	FirstName string             `json:"first_name" bson:"first_name"`
	LastName  string             `json:"last_name" bson:"last_name"`
	Username  string             `json:"username" bson:"username"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
}
