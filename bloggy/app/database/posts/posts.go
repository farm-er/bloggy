package posts

import (
	"context"
	"errors"

	"github.com/farm-er/bloggy/database"
	"go.mongodb.org/mongo-driver/bson"
)

func (p *Post) GetPost() error {

	err := database.OpenCollection("posts").FindOne(context.TODO(), bson.M{"_id": p.Id}).Decode(p)

	if err != nil {
		return err
	}

	return nil
}

func (p *Post) AddPost() error {

	_, err := database.OpenCollection("posts").InsertOne(context.TODO(), p)

	if err != nil {
		return err
	}

	return nil

}

func (p *Post) RemovePost() error {

	result, err := database.OpenCollection("posts").DeleteOne(context.TODO(), bson.M{"_id": p.Id})

	if err != nil {
		return err
	}

	if result.DeletedCount != 1 {
		return errors.New("error deleting post")
	}
	return nil
}

func (p *Post) LikePost() error {

	result, err := database.OpenCollection("posts").UpdateByID(context.TODO(), p.Id, bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "likes",
					Value: p.Likes + 1},
			}}})

	if err != nil {
		return err
	}

	if result.ModifiedCount != 1 {
		return errors.New("error updating post")
	}

	return nil
}

func (p *Post) CommentPost() error {

	result, err := database.OpenCollection("posts").UpdateByID(context.TODO(), p.Id, bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "number_comments",
					Value: p.NumberComments + 1},
			}}})

	if err != nil {
		return err
	}

	if result.ModifiedCount != 1 {
		return errors.New("error updating post")
	}

	return nil
}

func (p *Post) UnCommentPost() error {

	result, err := database.OpenCollection("posts").UpdateByID(context.TODO(), p.Id, bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "number_comments",
					Value: p.NumberComments - 1},
			}}})

	if err != nil {
		return err
	}

	if result.ModifiedCount != 1 {
		return errors.New("error updating post")
	}

	return nil
}
