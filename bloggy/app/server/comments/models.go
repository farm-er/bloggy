package comments

type AddCommentData struct {
	PostId string `json:"post_id"`
	Body   string `json:"body"`
}

type DeleteCommentData struct {
	CommentId string `json:"comment_id"`
	PostId    string `json:"post_id"`
}
