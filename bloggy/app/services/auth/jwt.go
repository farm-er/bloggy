package auth

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func GenerateJWToken(payload TokenData) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	secretKey := os.Getenv("JWT_KEY")

	// new token
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(1000 * time.Minute).Unix()
	claims["authorized"] = true
	claims["email"] = payload.Email
	claims["id"] = payload.Id
	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// verify token and return uid if success and error if fail
func VerifyToken(header http.Header) (string, error) {

	token := header.Get("token")

	if token == "" {
		return "Couldn't find token", errors.New("couldn't find token")
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	key := os.Getenv("JWT_KEY")

	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return "error decoding token", err
	}

	// return id
	for k, v := range claims {
		if k == "id" {
			return v.(string), nil
		}
	}
	return "Error extracting data", errors.New("error extracting data")
}

func HashPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	if err != nil {
		return "Error hashing password", err
	}
	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
