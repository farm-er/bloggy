package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	listenAdr string
}

type ApiError struct {
	Error string
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func NewApiServer(listenadr string) *ApiServer {
	return &ApiServer{
		listenAdr: listenadr,
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func makeHTTPHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func (s *ApiServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/", makeHTTPHandlerFunc(rootHandler))
	http.ListenAndServe(s.listenAdr, router)
}

func rootHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}
