package store

import (
	"github.com/oussama4/commentify/model"
)

type Store interface {
	GetComment(id string) (*model.Comment, error)
	CreateComment(body, parentId, userId, threadId string) (string, error)
	UpdateComment(id string, body string) error
	DeleteComment(id string) error
	ListComments(threadId string, parentId string, page int, pageSize int) ([]*model.CommentOutput, error)
	GetThread(id string) (*model.Thread, error)
	CreateThread(url, domain, title string) (string, error)
	GetUser(id string) (*model.User, error)
	CreateUser(name, email string) (string, error)
}
