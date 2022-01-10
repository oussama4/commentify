package handlers

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/oussama4/commentify/store"
)

func Routes(store store.Store, logger *log.Logger) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "Content-Range", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	u := &User{
		store:  store,
		logger: logger,
	}
	th := &Page{
		store:  store,
		logger: logger,
	}
	c := &Comment{
		store:  store,
		logger: logger,
	}
	r.Mount("/users", u.routes())
	r.Mount("/pages", th.routes())
	r.Mount("/comments", c.routes())

	return r
}
