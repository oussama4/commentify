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
