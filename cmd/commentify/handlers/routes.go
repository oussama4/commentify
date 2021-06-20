package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/oussama4/commentify/store"
)

func Routes(store store.Store) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	u := &User{
		store: store,
	}
	th := &Thread{store: store}
	r.Mount("/users", u.routes())
	r.Mount("/threads", th.routes())

	return r
}
