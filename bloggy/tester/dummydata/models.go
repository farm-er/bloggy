package dummydata

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

type Post struct {
	Id             primitive.ObjectID `json:"id" bson:"_id"`
	OwnerId        primitive.ObjectID `json:"owner_id" bson:"owner_id"`
	Title          string             `json:"title" bson:"title"`
	Body           string             `json:"body" bson:"body"`
	Likes          int                `json:"likes" bson:"likes"`
	NumberComments int                `json:"number_comments" bson:"number_comments"`
}

type Comment struct {
	Id     primitive.ObjectID `json:"id" bson:"_id"`
	PostId primitive.ObjectID `json:"post_id" bson:"post_id"`
	UserId primitive.ObjectID `json:"user_id" bson:"user_id"`
	Body   primitive.ObjectID `json:"body" bson:"body"`
	Likes  int                `json:"likes" bson:"likes"`
}
