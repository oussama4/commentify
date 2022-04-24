package store

import "github.com/oussama4/commentify/business/data/model"

type Store interface {
	GetComment(id string) (*model.Comment, error)
	CreateComment(body, parentId, userId, pageId string) (string, error)
	UpdateComment(id string, body string) error
	DeleteComment(id string) error
	ListComments(pageId string, pageUrl string, parentId string, page int, pageSize int) ([]model.CommentOutput, error)
	CountComments(pageId string) (int, error)
	GetPage(id string) (*model.Page, error)
	ListPages(page int, pageSize int) ([]model.Page, error)
	CreatePage(url, title string) (string, error)
	GetUser(id string) (*model.User, error)
	GetOrCreateGuest(name string) (string, error)
	ListUsers(page int, pageSize int) ([]model.User, error)
	CreateUser(name, email string) (string, error)
}
