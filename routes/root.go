package routes

import (
	"bloggy_api/database"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ApiServer struct {
	listenAdr   string
	mongoclient *mongo.Client
}

type ApiError struct {
	Error string
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func NewApiServer(listenadr string, mongoclient *mongo.Client) *ApiServer {
	return &ApiServer{
		listenAdr:   listenadr,
		mongoclient: mongoclient,
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

	router.HandleFunc("/", makeHTTPHandlerFunc(s.rootHandler))
	fmt.Printf("New connection on port: %v\n", s.listenAdr)
	http.ListenAndServe(s.listenAdr, router)

}

func (s *ApiServer) rootHandler(w http.ResponseWriter, r *http.Request) error {

	if r.Method == "GET" {
		return s.handleGetAllPosts(w, r)
	}
	return nil
}

func (s *ApiServer) handleGetAllPosts(w http.ResponseWriter, r *http.Request) error {

	collection := s.mongoclient.Database("bloggy").Collection("posts")
	filter := bson.D{{"title", "Welcome to BLOGGY"}}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var posts []database.BlogPost
	if err = cursor.All(context.TODO(), &posts); err != nil {
		panic(err)
	}
	for _, result := range posts {
		res, _ := bson.MarshalExtJSON(result, false, false)
		fmt.Println(string(res))
	}
	return WriteJSON(w, 200, posts)
}
