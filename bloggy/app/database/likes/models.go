package likes

import "go.mongodb.org/mongo-driver/bson/primitive"

type Like struct {
	PostId primitive.ObjectID `json:"post_id" bson:"post_id"`
	UserId primitive.ObjectID `json:"user_id" bson:"user_id"`
}

type CommentLike struct {
	CommentID primitive.ObjectID `json:"comment_id" bson:"comment_id"`
	UserId    primitive.ObjectID `json:"user_id" bson:"user_id"`
}
