package request

type PostComment struct {
	PostID  string `json:"postId"`
	Comment string `json:"comment"`
}
