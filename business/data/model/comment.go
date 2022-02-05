package model

type Comment struct {
	Id        string `json:"id"`
	Body      string `json:"body"`
	ParentId  string `json:"parentId"`
	UserId    string `json:"userId"`
	PageId    string `json:"pageId"`
	CreatedAt string `json:"createdAt"`
}

type CommentOutput struct {
	Id        string `json:"id"`
	Body      string `json:"body"`
	CreatedAt string `json:"createdAt"`
	Author    User   `json:"user"`
}
