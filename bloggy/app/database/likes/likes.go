package likes

import (
	"context"
	"errors"

	"github.com/farm-er/bloggy/database"
	"go.mongodb.org/mongo-driver/bson"
)

func (l *Like) AddLike() error {

	_, err := database.OpenCollection("likes").InsertOne(context.TODO(), l)

	if err != nil {
		return err
	}

	return nil
}

func (l *Like) CheckLikedPosts() error {

	err := database.OpenCollection("likes").FindOne(context.TODO(), bson.M{"post_id": l.PostId, "user_id": l.UserId}).Decode(l)

	if err != nil {
		return nil
	}

	return errors.New("user already liked this post")
}

func (cl *CommentLike) AddLike() error {

	_, err := database.OpenCollection("comment_likes").InsertOne(context.TODO(), cl)

	if err != nil {
		return err
	}

	return nil
}

func (cl *CommentLike) CheckLikedcomments() error {

	err := database.OpenCollection("comment_likes").FindOne(context.TODO(), bson.M{"comment_id": cl.CommentID, "user_id": cl.UserId}).Decode(cl)

	if err != nil {
		return nil
	}

	return errors.New("user already liked this comment")
}
