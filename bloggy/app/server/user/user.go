package user

import (
	"encoding/json"
	"net/http"

	"github.com/farm-er/bloggy/database/user"
	"github.com/farm-er/bloggy/services"
	"github.com/farm-er/bloggy/services/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUser(w http.ResponseWriter, r *http.Request) error {

	id, err := auth.VerifyToken(r.Header)

	if err != nil {
		return err
	}
	var user user.User
	user.Id, err = primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	err = user.GetUserById()

	if err != nil {
		return err
	}

	err = services.WriteJson(w, http.StatusOK, user)

	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(w http.ResponseWriter, r *http.Request) error {

	id, err := auth.VerifyToken(r.Header)

	if err != nil {
		return err
	}

	var data UpdateUserData

	err = json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		return err
	}

	var newUser user.User
	newUser.Id, err = primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}
	err = newUser.GetUserById()

	if err != nil {
		return err
	}

	err = auth.VerifyPassword(newUser.Password, data.Password)

	if err != nil {
		return err
	}

	newUser.Username = data.Username

	err = newUser.UpdateUserUsername()

	if err != nil {
		return err
	}

	err = services.WriteJson(w, http.StatusOK, newUser)

	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(w http.ResponseWriter, r *http.Request) error {

	id, err := auth.VerifyToken(r.Header)

	if err != nil {
		return err
	}

	var data DeleteUserData
	err = json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		return err
	}

	var user user.User
	user.Id, err = primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	err = user.GetUserById()

	if err != nil {
		return err
	}

	err = auth.VerifyPassword(user.Password, data.Password)

	if err != nil {
		return err
	}

	err = user.RemoveUser()

	if err != nil {
		return err
	}

	err = services.WriteJson(w, http.StatusOK, "user removed")

	if err != nil {
		return err
	}

	return nil
}
