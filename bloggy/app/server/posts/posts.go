package posts

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/farm-er/bloggy/database/comments"
	"github.com/farm-er/bloggy/database/posts"
	"github.com/farm-er/bloggy/database/user"
	"github.com/farm-er/bloggy/services"
	"github.com/farm-er/bloggy/services/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddPost(w http.ResponseWriter, r *http.Request) error {

	id, err := auth.VerifyToken(r.Header)

	if err != nil {
		return err
	}

	var data AddPostData
	err = json.NewDecoder(r.Body).Decode(&data)

	if data.Password == "" || data.Title == "" || data.Body == "" {
		return errors.New("invalid data")
	}

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

	post := posts.Post{
		CreatedAt:      primitive.NewDateTimeFromTime(time.Now()),
		Id:             primitive.NewObjectID(),
		OwnerId:        newUser.Id,
		Title:          data.Title,
		Body:           data.Body,
		Likes:          0,
		NumberComments: 0,
	}

	err = post.AddPost()

	if err != nil {
		return err
	}

	err = services.WriteJson(w, http.StatusCreated, "post added")

	if err != nil {
		return err
	}

	return nil
}

func DeletePost(w http.ResponseWriter, r *http.Request) error {

	id, err := auth.VerifyToken(r.Header)

	if err != nil {
		return err
	}

	var data DeletePostData
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

	var post posts.Post
	post.Id, err = primitive.ObjectIDFromHex(data.PostId)

	if err != nil {
		return err
	}

	err = post.RemovePost()

	if err != nil {
		return err
	}

	var comment comments.Comment
	comment.PostId = post.Id

	err = comment.DeleteComments()

	if err != nil {
		return err
	}

	err = services.WriteJson(w, http.StatusCreated, "post deleted")

	if err != nil {
		return err
	}

	return nil
}
