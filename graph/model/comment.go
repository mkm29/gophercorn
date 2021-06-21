package model

type Comment struct {
	ID     string `json:"id"`
	Body   string `json:"body"`
	PostID string `json:"postId"`
	UserID string `json:"userId"`
}
