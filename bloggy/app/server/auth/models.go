package auth

import "go.mongodb.org/mongo-driver/bson/primitive"

type SignupData struct {
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Username  string `json:"username" bson:"username"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
}

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseUser struct {
	CreatedAt primitive.DateTime `json:"created_at"`
	Id        primitive.ObjectID `json:"id"`
	FirstName string             `json:"first_name"`
	LastName  string             `json:"last_name"`
	Username  string             `json:"username"`
	Email     string             `json:"email"`
}

type ResponseData struct {
	Token string       `json:"token"`
	User  ResponseUser `json:"user"`
}
