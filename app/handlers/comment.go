package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oussama4/commentify/base/web"
	"github.com/oussama4/commentify/store"
)

type Comment struct {
	store  store.Store
	logger *log.Logger
}

type CommentInput struct {
	ParentId string `json:"parent_id"`
	UserId   string `json:"user_id"`
	PageId   string `json:"page_id"`
	Body     string `json:"body"`
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
		respondError(c.logger, w, http.StatusInternalServerError, err.Error())
	}

	commentId, err := c.store.CreateComment(comment.Body, comment.ParentId, comment.UserId, comment.PageId)
	if err != nil {
		respondError(c.logger, w, http.StatusInternalServerError, err.Error())
	}

	web.Json(w, http.StatusOK, map[string]interface{}{"comment_id": commentId})
}

func (c *Comment) Get(w http.ResponseWriter, r *http.Request) {
	commentId := chi.URLParam(r, "commentId")
	comment, err := c.store.GetComment(commentId)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			respondError(c.logger, w, http.StatusNotFound, err.Error())
		}
		respondError(c.logger, w, http.StatusInternalServerError, err.Error())
	}

	web.Json(w, http.StatusOK, map[string]interface{}{"comment": comment})
}

func (c *Comment) List(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query()
	pageId := web.ReadString(qs, "page_id", "")
	pageUrl := web.ReadString(qs, "page_url", "")
	parentId := web.ReadString(qs, "parent", "")
	page, _ := web.ReadInt(qs, "page", 0)
	pageSize, _ := web.ReadInt(qs, "page_size", 0)

	comments, err := c.store.ListComments(pageId, pageUrl, parentId, page, pageSize)
	if err != nil {
		respondError(c.logger, w, http.StatusInternalServerError, err.Error())
	}
	web.Json(w, http.StatusOK, map[string]interface{}{"comments": comments})
}

func (c *Comment) Count(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query()
	pageId := web.ReadString(qs, "page_id", "")
	count, err := c.store.CountComments(pageId)
	if err != nil {
		respondError(c.logger, w, http.StatusInternalServerError, err.Error())
	}

	web.Json(w, http.StatusOK, map[string]interface{}{"count": count})
}

func (c *Comment) Delete(w http.ResponseWriter, r *http.Request) {
	commentId := chi.URLParam(r, "commentId")
	if err := c.store.DeleteComment(commentId); err != nil {
		respondError(c.logger, w, http.StatusInternalServerError, err.Error())
	}

	web.Json(w, http.StatusOK, map[string]interface{}{"deleted": true})
}
