package model

type Comment struct {
	Id        string
	Body      string
	ParentId  string
	UserId    string
	ThreadId  string
	CreatedAt string
}

type CommentOutput struct {
	Id        string `json:"id"`
	Body      string `json:"body"`
	CreatedAt string `json:"createAt"`
	Author    User   `json:"user"`
}
