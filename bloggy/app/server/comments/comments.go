package comments

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

func AddComment(w http.ResponseWriter, r *http.Request) error {

	id, err := auth.VerifyToken(r.Header)

	if err != nil {
		return err
	}

	var data AddCommentData
	err = json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		return err
	}

	if data.Body == "" || data.PostId == "" {
		return errors.New("invalid data")
	}

	var newUser user.User
	newUser.Id, err = primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	// if user exists
	err = newUser.GetUserById()

	if err != nil {
		return err
	}

	// post exists
	var post posts.Post
	post.Id, err = primitive.ObjectIDFromHex(data.PostId)

	if err != nil {
		return err
	}

	// post exists
	err = post.GetPost()

	if err != nil {
		return err
	}
	comment := comments.Comment{
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		Id:        primitive.NewObjectID(),
		PostId:    post.Id,
		UserId:    newUser.Id,
		Body:      data.Body,
		Likes:     0,
	}

	err = comment.AddComment()

	if err != nil {
		return err
	}

	err = post.CommentPost()

	if err != nil {
		return err
	}

	// increment comments count

	err = services.WriteJson(w, http.StatusCreated, "comment added")

	if err != nil {
		return err
	}

	return nil
}

func DeleteComment(w http.ResponseWriter, r *http.Request) error {

	id, err := auth.VerifyToken(r.Header)

	if err != nil {
		return err
	}

	var data DeleteCommentData
	err = json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		return err
	}

	if data.CommentId == "" {
		return errors.New("invalid data")
	}

	var newUser user.User
	newUser.Id, err = primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	// if user exists
	err = newUser.GetUserById()

	if err != nil {
		return err
	}

	var comment comments.Comment
	comment.Id, err = primitive.ObjectIDFromHex(data.CommentId)

	if err != nil {
		return err
	}

	err = comment.GetComment()

	if err != nil {
		return err
	}

	if comment.UserId != newUser.Id {
		return errors.New("trying to remove another user's comment")
	}

	err = comment.RemoveComment()

	if err != nil {
		return err
	}

	var post posts.Post
	post.Id, err = primitive.ObjectIDFromHex(data.PostId)

	if err != nil {
		return err
	}

	err = post.GetPost()

	if err != nil {
		return err
	}
	err = post.UnCommentPost()

	if err != nil {
		return err
	}

	err = services.WriteJson(w, http.StatusCreated, "comment deleted")

	if err != nil {
		return err
	}

	return nil
}
