package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oussama4/commentify/base/validate"
	"github.com/oussama4/commentify/base/web"
	"github.com/oussama4/commentify/business/data/store"
)

type Page struct {
	store  store.Store
	logger *log.Logger
}

type PageInput struct {
	Url, Title string
}

func (th *Page) routes() http.Handler {
	r := chi.NewRouter()
	r.Post("/", th.Create)
	r.Get("/{pageId}", th.Get)
	r.Get("/", th.List)

	return r
}

func (pi *PageInput) valid() error {
	v := validate.New()

	v.Check(pi.Title != "", "title", "page title is required")
	v.Check(pi.Url != "", "url", "page url is required")

	return v.Valid()
}

func (p *Page) List(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query()
	page, _ := web.ReadInt(qs, "page", 0)
	pageSize, _ := web.ReadInt(qs, "page_size", 10)

	pages, err := p.store.ListPages(page, pageSize)
	if err != nil {
		respondError(p.logger, w, http.StatusInternalServerError, err)
	}
	web.Json(w, http.StatusOK, map[string]interface{}{"pages": pages})
}

func (p *Page) Create(w http.ResponseWriter, r *http.Request) {
	page := PageInput{}
	if err := json.NewDecoder(r.Body).Decode(&page); err != nil {
		respondError(p.logger, w, http.StatusInternalServerError, err)
	}

	if err := page.valid(); err != nil {
		respondError(p.logger, w, http.StatusUnprocessableEntity, err)
	}

	pageId, err := p.store.CreatePage(page.Url, page.Title)
	if err != nil {
		respondError(p.logger, w, http.StatusInternalServerError, err)
	}

	web.Json(w, http.StatusOK, map[string]interface{}{"page_id": pageId})
}

func (p *Page) Get(w http.ResponseWriter, r *http.Request) {
	pageId := chi.URLParam(r, "pageId")
	page, err := p.store.GetPage(pageId)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			respondError(p.logger, w, http.StatusNotFound, err)
		}
		respondError(p.logger, w, http.StatusInternalServerError, err)
	}

	web.Json(w, http.StatusOK, map[string]interface{}{"page": page})
}
