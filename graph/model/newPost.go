package model

type NewPost struct {
	Title  string `json:"text"`
	Body   string `json:"text"`
	UserID string `json:"userId"`
}
