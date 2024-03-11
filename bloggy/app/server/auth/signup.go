package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/mail"
	"sync"

	"github.com/farm-er/bloggy/database/user"
	"github.com/farm-er/bloggy/services"
	"github.com/farm-er/bloggy/services/auth"
)

func SignUp(w http.ResponseWriter, r *http.Request) error {

	var data SignupData
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		return err
	}

	if data.Email == "" || data.Password == "" || data.FirstName == "" || data.LastName == "" {
		return errors.New("invalid data")
	}

	_, err = mail.ParseAddress(data.Email)

	if err != nil {
		return err
	}

	data.Password, err = auth.HashPassword(data.Password)

	if err != nil {
		return err
	}

	user := user.User{
		Email:     data.Email,
		Username:  data.Username,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Password:  data.Password,
	}

	err = user.GetUserByEmail()

	if err == nil {
		return errors.New("email already exists")
	}

	err = user.GetUserByUsername()

	if err == nil {
		return errors.New("username already exists")
	}

	err = user.AddUser()

	if err != nil {
		return err
	}

	token, err := auth.GenerateJWToken(auth.TokenData{
		Id:    user.Id.Hex(),
		Email: user.Email,
	})

	if err != nil {
		return err
	}

	err = services.WriteJson(w, http.StatusOK, ResponseData{
		Token: token,
		User: ResponseUser{
			CreatedAt: user.CreatedAt,
			Id:        user.Id,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Username:  user.Username,
			Email:     user.Email,
		},
	})

	if err != nil {
		return err
	}

	return nil
}

func SyncSignUp(w http.ResponseWriter, r *http.Request) error {

	var data SignupData
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		return err
	}

	if data.Email == "" || data.Password == "" || data.FirstName == "" || data.LastName == "" {
		return errors.New("invalid data")
	}

	_, err = mail.ParseAddress(data.Email)

	if err != nil {
		return err
	}

	data.Password, err = auth.HashPassword(data.Password)

	if err != nil {
		return err
	}

	user := user.User{
		Email:     data.Email,
		Username:  data.Username,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Password:  data.Password,
	}

	errorCollector := make(chan error, 2)

	wg := &sync.WaitGroup{}

	wg.Add(2)

	go user.SyncGetUserByEmail(errorCollector, wg)
	go user.SyncGetUserByUsername(errorCollector, wg)

	wg.Wait()

	close(errorCollector)

	for err := range errorCollector {
		if err != nil {
			return err
		}
	}

	for err := range errorCollector {
		return err
	}
	err = user.AddUser()

	if err != nil {
		return err
	}

	token, err := auth.GenerateJWToken(auth.TokenData{
		Id:    user.Id.Hex(),
		Email: user.Email,
	})

	if err != nil {
		return err
	}

	err = services.WriteJson(w, http.StatusOK, ResponseData{
		Token: token,
		User: ResponseUser{
			CreatedAt: user.CreatedAt,
			Id:        user.Id,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Username:  user.Username,
			Email:     user.Email,
		},
	})

	if err != nil {
		return err
	}

	return nil
}
