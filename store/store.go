package store

import "github.com/oussama4/commentify/model"

type Store interface {
	GetComment(id string) (*model.Comment, error)
	CreateComment(comment model.Comment) (*model.Comment, error)
	UpdateComment(id string, body string) (*model.Comment, error)
	DeleteComment(id string) error
	ListComments(threadUrl string) ([]*model.Comment, error)
	GetThread(id string) (*model.Thread, error)
	CreateThread(thread model.Thread) (*model.Thread, error)
	CreateUser(user model.User) (*model.User, error)
}
