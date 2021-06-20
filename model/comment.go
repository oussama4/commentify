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
	Id        string
	Body      string
	CreatedAt string
	Author    User
}
