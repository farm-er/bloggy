package comments

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comment struct {
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"`
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	PostId    primitive.ObjectID `json:"post_id" bson:"post_id"`
	UserId    primitive.ObjectID `json:"user_id" bson:"user_id"`
	Body      string             `json:"body" bson:"body"`
	Likes     int                `json:"likes" bson:"likes"`
}
