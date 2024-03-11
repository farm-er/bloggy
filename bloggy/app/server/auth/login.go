package auth

import (
	"encoding/json"
	"net/http"

	"github.com/farm-er/bloggy/database/user"
	"github.com/farm-er/bloggy/services"
	"github.com/farm-er/bloggy/services/auth"
)

func Login(w http.ResponseWriter, r *http.Request) error {

	var data LoginData

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		return err
	}

	user := user.User{
		Email: data.Email,
	}

	err = user.GetUserByEmail()

	if err != nil {
		return err
	}

	err = auth.VerifyPassword(user.Password, data.Password)

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

	services.WriteJson(w, http.StatusOK, ResponseData{
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

	return nil
}
