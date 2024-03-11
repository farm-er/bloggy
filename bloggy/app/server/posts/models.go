package posts

type AddPostData struct {
	Title    string `json:"title"`
	Body     string `json:"body"`
	Password string `json:"password"`
}

type DeletePostData struct {
	PostId   string `json:"post_id"`
	Password string `json:"password"`
}
