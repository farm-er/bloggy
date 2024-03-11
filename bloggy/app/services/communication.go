package services

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, status int, v any) error {

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(v)

	if err != nil {
		return err
	}
	w.WriteHeader(status)
	return nil
}
