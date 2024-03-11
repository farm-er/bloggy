package dummydata

import (
	"fmt"
	"math/rand/v2"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var firstNames []string = []string{
	"oussama",
	"rayane",
	"simo",
	"aymmen",
	"salma",
	"kawtar",
	"ahmed",
	"amine",
	"youssef",
	"youness",
	"ali",
	"reda",
}

var lastNames []string = []string{
	"oussama",
	"rayane",
	"simo",
	"aymmen",
	"salma",
	"kawtar",
	"ahmed",
	"amine",
	"youssef",
	"youness",
	"ali",
	"reda",
}

func UserGenerator() *User {
	firstname := firstNames[rand.IntN(11)] + fmt.Sprint(rand.IntN(685712))
	return &User{
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		Id:        primitive.NewObjectID(),
		FirstName: firstname,
		LastName:  lastNames[rand.IntN(11)] + fmt.Sprint(rand.IntN(685712)),
		Username:  firstname,
		Email:     firstname + "@gmail.com",
		Password:  "anything",
	}
}
