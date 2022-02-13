package model

import "github.com/oussama4/commentify/base/validate"

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
	UserId   string `json:"user_id"`
	PageId   string `json:"page_id"`
	Body     string `json:"body"`
}

type CommentOutput struct {
	Id        string `json:"id"`
	Body      string `json:"body"`
	CreatedAt string `json:"createdAt"`
	Author    User   `json:"user"`
}

func (ci *CommentInput) Valid() error {
	v := validate.New()

	v.Check(ci.UserId != "", "user_id", "user_id is required")
	v.Check(ci.PageId != "", "page_id", "page_id is required")
	v.Check(ci.Body != "", "body", "you didn't write a comment")

	return v.Valid()
}
