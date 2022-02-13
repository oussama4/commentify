package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oussama4/commentify/base/validate"
	"github.com/oussama4/commentify/base/web"
	"github.com/oussama4/commentify/business/data/model"
	"github.com/oussama4/commentify/business/data/store"
)

type Comment struct {
	store  store.Store
	logger *log.Logger
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
	comment := model.CommentInput{}
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		respondError(c.logger, w, http.StatusInternalServerError, err)
	}

	if err := comment.Valid(); err != nil {
		respondError(c.logger, w, http.StatusUnprocessableEntity, err)
	}

	commentId, err := c.store.CreateComment(comment.Body, comment.ParentId, comment.UserId, comment.PageId)
	if err != nil {
		respondError(c.logger, w, http.StatusInternalServerError, err)
	}

	web.Json(w, http.StatusOK, map[string]interface{}{"comment_id": commentId})
}

func (c *Comment) Get(w http.ResponseWriter, r *http.Request) {
	commentId := chi.URLParam(r, "commentId")
	comment, err := c.store.GetComment(commentId)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			respondError(c.logger, w, http.StatusNotFound, err)
		}
		respondError(c.logger, w, http.StatusInternalServerError, err)
	}

	web.Json(w, http.StatusOK, map[string]interface{}{"comment": comment})
}

func (c *Comment) List(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query()
	pageId := web.ReadString(qs, "page_id", "")
	pageUrl := web.ReadString(qs, "page_url", "")
	parentId := web.ReadString(qs, "parent", "")
	page, _ := web.ReadInt(qs, "page", 0)
	pageSize, _ := web.ReadInt(qs, "page_size", 10)

	v := validate.New()
	v.Check(pageId != "" || pageUrl != "", "page_id", "page_id or page_url must be provided")
	if err := v.Valid(); err != nil {
		respondError(c.logger, w, http.StatusUnprocessableEntity, err)
	}

	comments, err := c.store.ListComments(pageId, pageUrl, parentId, page, pageSize)
	if err != nil {
		respondError(c.logger, w, http.StatusInternalServerError, err)
	}

	web.Json(w, http.StatusOK, map[string]interface{}{"comments": comments})
}

func (c *Comment) Count(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query()
	pageId := web.ReadString(qs, "page_id", "")
	count, err := c.store.CountComments(pageId)
	if err != nil {
		respondError(c.logger, w, http.StatusInternalServerError, err)
	}

	web.Json(w, http.StatusOK, map[string]interface{}{"count": count})
}

func (c *Comment) Delete(w http.ResponseWriter, r *http.Request) {
	commentId := chi.URLParam(r, "commentId")
	if err := c.store.DeleteComment(commentId); err != nil {
		respondError(c.logger, w, http.StatusInternalServerError, err)
	}

	web.Json(w, http.StatusOK, map[string]interface{}{"deleted": true})
}
