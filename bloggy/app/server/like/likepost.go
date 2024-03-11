package like

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/farm-er/bloggy/database/likes"
	"github.com/farm-er/bloggy/database/posts"
	"github.com/farm-er/bloggy/database/user"
	"github.com/farm-er/bloggy/services"
	"github.com/farm-er/bloggy/services/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func LikePost(w http.ResponseWriter, r *http.Request) error {

	id, err := auth.VerifyToken(r.Header)

	if err != nil {
		return err
	}

	var data LikePostData
	err = json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		return err
	}

	if data.PostId == "" {
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

	var post posts.Post
	post.Id, err = primitive.ObjectIDFromHex(data.PostId)

	if err != nil {
		return err
	}

	err = post.GetPost()

	if err != nil {
		return err
	}

	like := likes.Like{
		PostId: post.Id,
		UserId: newUser.Id,
	}

	// need to check if the post is already liked by this user
	err = like.CheckLikedPosts()

	if err != nil {
		return err
	}

	err = post.LikePost()

	if err != nil {
		return err
	}

	err = like.AddLike()

	if err != nil {
		return err
	}

	err = services.WriteJson(w, http.StatusCreated, "post liked")

	if err != nil {
		return err
	}

	return nil
}
