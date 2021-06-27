package handlers

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/oussama4/commentify/store"
)

func Routes(store store.Store, logger *log.Logger) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	u := &User{
		store:  store,
		logger: logger,
	}
	th := &Thread{
		store:  store,
		logger: logger,
	}
	c := &Comment{
		store:  store,
		logger: logger,
	}
	r.Mount("/users", u.routes())
	r.Mount("/threads", th.routes())
	r.Mount("/comments", c.routes())

	return r
}
