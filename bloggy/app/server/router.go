package server

import (
	"net/http"

	"github.com/farm-er/bloggy/server/auth"
	"github.com/farm-er/bloggy/server/comments"
	"github.com/farm-er/bloggy/server/like"
	"github.com/farm-er/bloggy/server/posts"
	"github.com/farm-er/bloggy/server/user"
	"github.com/farm-er/bloggy/services"
)

func (a *Apiserver) handleCommentLike(w http.ResponseWriter, r *http.Request) {

	// like a post
	if r.Method == "POST" {
		err := like.LikeComment(w, r)
		if err != nil {
			services.WriteJson(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	services.WriteJson(w, http.StatusForbidden, "invalid method boy")
}

func (a *Apiserver) handlePostLike(w http.ResponseWriter, r *http.Request) {

	// like a post
	if r.Method == "POST" {
		err := like.LikePost(w, r)
		if err != nil {
			services.WriteJson(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	services.WriteJson(w, http.StatusForbidden, "invalid method boy")

}

func (a *Apiserver) handlecomments(w http.ResponseWriter, r *http.Request) {

	// add comment
	if r.Method == "POST" {
		err := comments.AddComment(w, r)
		if err != nil {
			services.WriteJson(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	// delete comment
	if r.Method == "DELETE" {
		err := comments.DeleteComment(w, r)
		if err != nil {
			services.WriteJson(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	services.WriteJson(w, http.StatusForbidden, "invalid method boy")

}

func (a *Apiserver) handlePosts(w http.ResponseWriter, r *http.Request) {

	// add post
	if r.Method == "POST" {
		err := posts.AddPost(w, r)
		if err != nil {
			services.WriteJson(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	// delete post
	if r.Method == "DELETE" {
		err := posts.DeletePost(w, r)
		if err != nil {
			services.WriteJson(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	services.WriteJson(w, http.StatusForbidden, "invalid method boy")

}

func (a *Apiserver) handleUser(w http.ResponseWriter, r *http.Request) {

	// get user account
	if r.Method == "GET" {
		err := user.GetUser(w, r)
		if err != nil {
			services.WriteJson(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	// Update user info
	if r.Method == "POST" {
		err := user.UpdateUser(w, r)
		if err != nil {
			services.WriteJson(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	// delete user account
	if r.Method == "DELETE" {
		err := user.DeleteUser(w, r)
		if err != nil {
			services.WriteJson(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	services.WriteJson(w, http.StatusForbidden, "invalid method boy")
}

func (a *Apiserver) handleSignup(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		err := auth.SyncSignUp(w, r)
		if err != nil {
			services.WriteJson(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	services.WriteJson(w, http.StatusBadGateway, "invalid method boy")
}

func (a *Apiserver) handleLogin(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		err := auth.Login(w, r)
		if err != nil {
			services.WriteJson(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	services.WriteJson(w, http.StatusBadGateway, "invalid method boy")
}
