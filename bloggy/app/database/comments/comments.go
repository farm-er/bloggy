package comments

import (
	"context"
	"errors"

	"github.com/farm-er/bloggy/database"
	"go.mongodb.org/mongo-driver/bson"
)

func (c *Comment) GetComment() error {

	err := database.OpenCollection("comments").FindOne(context.TODO(), bson.M{"_id": c.Id}).Decode(c)

	if err != nil {
		return err
	}

	return nil
}

func (c *Comment) AddComment() error {

	_, err := database.OpenCollection("comments").InsertOne(context.TODO(), c)

	if err != nil {
		return err
	}

	return nil
}

func (c *Comment) RemoveComment() error {

	result, err := database.OpenCollection("comments").DeleteOne(context.TODO(), bson.M{"_id": c.Id})

	if err != nil {
		return err
	}

	if result.DeletedCount != 1 {
		return errors.New("error deleting comment")
	}

	return nil
}

func (c *Comment) DeleteComments() error {

	_, err := database.OpenCollection("comments").DeleteMany(context.TODO(), bson.M{"post_id": c.PostId})

	if err != nil {
		return err
	}

	return nil
}

func (c *Comment) LikeComment() error {

	result, err := database.OpenCollection("comments").UpdateByID(context.TODO(), c.Id, bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "likes",
					Value: c.Likes + 1},
			}}})

	if err != nil {
		return err
	}

	if result.ModifiedCount != 1 {
		return errors.New("error updating comment")
	}

	return nil
}
