package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oussama4/commentify/store"
	"github.com/oussama4/commentify/web"
)

type Thread struct {
	store store.Store
}

type ThreadInput struct {
	Url, Domain, Title string
}

func (th *Thread) routes() http.Handler {
	r := chi.NewRouter()
	r.Post("/", th.Create)
	r.Get("/{threadId}", th.Get)

	return r
}

func (th *Thread) Create(w http.ResponseWriter, r *http.Request) {
	thread := ThreadInput{}
	if err := json.NewDecoder(r.Body).Decode(&thread); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	threadId, err := th.store.CreateThread(thread.Url, thread.Domain, thread.Title)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	web.Json(w, threadId)
}

func (th *Thread) Get(w http.ResponseWriter, r *http.Request) {
	threadId := chi.URLParam(r, "threadId")
	thread, err := th.store.GetThread(threadId)
	if err != nil {
		if err == store.ErrNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	web.Json(w, thread)
}
