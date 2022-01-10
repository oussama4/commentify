package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oussama4/commentify/store"
	"github.com/oussama4/commentify/web"
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

func (p *Page) List(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query()
	page, _ := web.ReadInt(qs, "page", 0)
	pageSize, _ := web.ReadInt(qs, "page_size", 10)

	pages, err := p.store.ListPages(page, pageSize)
	if err != nil {
		p.logger.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	web.Json(w, pages)
}

func (p *Page) Create(w http.ResponseWriter, r *http.Request) {
	page := PageInput{}
	if err := json.NewDecoder(r.Body).Decode(&page); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	pageId, err := p.store.CreatePage(page.Url, page.Title)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	web.Json(w, pageId)
}

func (p *Page) Get(w http.ResponseWriter, r *http.Request) {
	pageId := chi.URLParam(r, "pageId")
	page, err := p.store.GetPage(pageId)
	if err != nil {
		if err == store.ErrNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	web.Json(w, page)
}
