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
	ParentId string
	UserId   string
	PageId   string
	Body     string
}

func (c *Comment) routes() http.Handler {
	r := chi.NewRouter()
	r.Post("/", c.Create)
	r.Get("/{commentId}", c.Get)
	r.Get("/", c.List)
	r.Get("/count", c.Count)
	r.Delete("/{commentId}", c.Delete)

	return r
}

func (c *Comment) Create(w http.ResponseWriter, r *http.Request) {
	comment := CommentInput{}
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	commentId, err := c.store.CreateComment(comment.Body, comment.ParentId, comment.UserId, comment.PageId)
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
	qs := r.URL.Query()
	pageId := web.ReadString(qs, "page_id", "")
	parentId := web.ReadString(qs, "parent", "")
	page, _ := web.ReadInt(qs, "page", 0)
	pageSize, _ := web.ReadInt(qs, "page_size", 0)

	comments, err := c.store.ListComments(pageId, parentId, page, pageSize)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	web.Json(w, comments)
}

func (c *Comment) Count(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query()
	pageId := web.ReadString(qs, "page_id", "")
	count, err := c.store.CountComments(pageId)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	web.Json(w, map[string]int{"count": count})
}

func (c *Comment) Delete(w http.ResponseWriter, r *http.Request) {
	commentId := chi.URLParam(r, "commentId")
	if err := c.store.DeleteComment(commentId); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	web.Json(w, map[string]bool{"deleted": true})
}
