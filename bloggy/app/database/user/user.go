package user

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/farm-er/bloggy/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *User) GetUserByEmail() error {
	err := database.OpenCollection("users").FindOne(context.TODO(), bson.M{"email": u.Email}).Decode(u)

	if err != nil {
		return err
	}

	return nil
}

func (u *User) SyncGetUserByEmail(errorCollector chan error, wg *sync.WaitGroup) {

	defer wg.Done()

	err := database.OpenCollection("users").FindOne(context.TODO(), bson.M{"email": u.Email}).Decode(u)

	if err == nil {
		errorCollector <- errors.New("email already exists")
	}
}

func (u *User) GetUserByUsername() error {

	err := database.OpenCollection("users").FindOne(context.TODO(), bson.M{"username": u.Username}).Decode(u)

	if err != nil {
		return err
	}
	return nil
}

func (u *User) SyncGetUserByUsername(errorCollector chan error, wg *sync.WaitGroup) {

	defer wg.Done()

	err := database.OpenCollection("users").FindOne(context.TODO(), bson.M{"username": u.Username}).Decode(u)

	if err == nil || u.Id != primitive.NilObjectID {
		errorCollector <- errors.New("username already exists")
	}
}

func (u *User) GetUserById() error {

	err := database.OpenCollection("users").FindOne(context.TODO(), bson.M{"_id": u.Id}).Decode(u)

	if err != nil {
		return err
	}

	return nil
}

func (u *User) AddUser() error {

	u.Id = primitive.NewObjectID()
	u.CreatedAt = primitive.NewDateTimeFromTime(time.Now())

	_, err := database.OpenCollection("users").InsertOne(context.TODO(), u)

	if err != nil {
		return err
	}
	return nil
}

func (u *User) UpdateUserUsername() error {

	result, err := database.OpenCollection("users").UpdateByID(context.TODO(), u.Id, bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "username",
					Value: u.Username},
			}}})

	if err != nil {
		return err
	}

	if result.ModifiedCount != 1 {
		return errors.New("error replacing user")
	}

	return nil
}

func (u *User) ReplaceUser() error {

	result, err := database.OpenCollection("users").ReplaceOne(context.TODO(), bson.M{"_id": u.Id}, u)

	if err != nil {
		return err
	}

	if result.ModifiedCount != 1 {
		return errors.New("error replacing user")
	}

	return nil
}

func (u *User) RemoveUser() error {

	result, err := database.OpenCollection("users").DeleteOne(context.TODO(), bson.M{"_id": u.Id})

	if err != nil {
		return err
	}

	if result.DeletedCount != 1 {
		return errors.New("error deleting user")
	}

	return nil

}
