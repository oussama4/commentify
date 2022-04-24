package model

import (
	"strings"

	"github.com/oussama4/commentify/base/validate"
)

type Comment struct {
	Id        string `json:"id"`
	Body      string `json:"body"`
	ParentId  string `json:"parentId"`
	UserId    string `json:"userId"`
	PageId    string `json:"pageId"`
	CreatedAt string `json:"createdAt"`
}
type CommentInput struct {
	ParentId string `json:"parent_id"`
	Guest    string `jspn:"guest"`
	UserId   string `json:"user_id"`
	PageId   string `json:"page_id"`
	Body     string `json:"body"`
}

type CommentOutput struct {
	Id        string `json:"id"`
	Body      string `json:"body"`
	CreatedAt string `json:"createdAt"`
	Author    User   `json:"author"`
}

func (ci *CommentInput) Valid() error {
	v := validate.New()

	v.Check(ci.UserId != "" && strings.TrimSpace(ci.Guest) == "", "user_id", "user_id is required")
	v.Check(ci.PageId != "", "page_id", "page_id is required")
	v.Check(strings.TrimSpace(ci.Body) != "", "body", "you didn't write a comment")

	return v.Valid()
}
