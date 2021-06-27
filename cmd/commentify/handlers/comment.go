package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oussama4/commentify/store"
	"github.com/oussama4/commentify/web"
)

type Comment struct {
	store  store.Store
	logger *log.Logger
}

type CommentInput struct {
	Body     string
	ParentId string
	UserId   string
	ThreadId string
}

func (c *Comment) routes() http.Handler {
	r := chi.NewRouter()
	r.Post("/", c.Create)
	r.Get("/{commentId}", c.Get)
	r.Get("/thread/{threadId}", c.List)
	r.Get("/thread/{threadId}/{parentId}", c.List)

	return r
}

func (c *Comment) Create(w http.ResponseWriter, r *http.Request) {
	comment := CommentInput{}
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	commentId, err := c.store.CreateComment(comment.Body, comment.ParentId, comment.UserId, comment.ThreadId)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	web.Json(w, commentId)
}

func (c *Comment) Get(w http.ResponseWriter, r *http.Request) {
	commentId := chi.URLParam(r, "commentId")
	comment, err := c.store.GetComment(commentId)
	if err != nil {
		if err == store.ErrNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	web.Json(w, comment)
}

func (c *Comment) List(w http.ResponseWriter, r *http.Request) {
	threadId := chi.URLParam(r, "threadId")
	parentId := chi.URLParam(r, "parentId")
	if parentId == "" {
		comments, err := c.store.ListComments(threadId)
		if err != nil {
			c.logger.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		web.Json(w, comments)
	} else {
		comments, err := c.store.ListChildComments(threadId, parentId)
		if err != nil {
			c.logger.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		web.Json(w, comments)
	}
}
