package model

import "time"

type Comment struct {
	Id        string
	Body      string
	ParentId  string
	UserId    string
	ThreadId  string
	CreatedAt time.Time
}

type CommentOutput struct {
	Id        string
	Body      string
	CreatedAt time.Time
	Author    User
}
