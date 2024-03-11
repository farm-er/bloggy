package server

import (
	"fmt"
	"log"
	"net/http"
)

type Apiserver struct {
	Addr string
}

func CreateServer(addr string) *Apiserver {

	return &Apiserver{
		Addr: addr,
	}

}

func (a *Apiserver) Run() {

	http.HandleFunc("/user", a.handleUser)
	http.HandleFunc("/signup", a.handleSignup)
	http.HandleFunc("/login", a.handleLogin)
	http.HandleFunc("/posts", a.handlePosts)
	http.HandleFunc("/posts/like", a.handlePostLike)
	http.HandleFunc("/comments", a.handlecomments)
	http.HandleFunc("/comments/like", a.handleCommentLike)
	fmt.Println("server is running on port", a.Addr)

	//err := http.ListenAndServeTLS(a.Addr, "server.crt", "server.key", nil)

	err := http.ListenAndServe(a.Addr, nil)

	if err != nil {
		log.Fatal(err)
	}
}
