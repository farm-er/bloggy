package posts

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	CreatedAt      primitive.DateTime `json:"created_at" bson:"created_at"`
	Id             primitive.ObjectID `json:"id" bson:"_id"`
	OwnerId        primitive.ObjectID `json:"owner_id" bson:"owner_id"`
	Title          string             `json:"title" bson:"title"`
	Body           string             `json:"body" bson:"body"`
	Likes          int                `json:"likes" bson:"likes"`
	NumberComments int                `json:"number_comments" bson:"number_comments"`
}
