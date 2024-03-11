package like

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/farm-er/bloggy/database/comments"
	"github.com/farm-er/bloggy/database/likes"
	"github.com/farm-er/bloggy/database/user"
	"github.com/farm-er/bloggy/services"
	"github.com/farm-er/bloggy/services/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func LikeComment(w http.ResponseWriter, r *http.Request) error {

	id, err := auth.VerifyToken(r.Header)

	if err != nil {
		return err
	}

	var data LikeCommentData
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

	like := likes.CommentLike{
		CommentID: comment.Id,
		UserId:    newUser.Id,
	}

	// need to check if the comment is already liked by this user

	err = like.CheckLikedcomments()

	if err != nil {
		return err
	}

	err = comment.LikeComment()

	if err != nil {
		return err
	}

	err = like.AddLike()

	if err != nil {
		return err
	}

	err = services.WriteJson(w, http.StatusCreated, "comment liked")

	if err != nil {
		return err
	}

	return nil
}
