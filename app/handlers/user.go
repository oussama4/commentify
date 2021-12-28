package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oussama4/commentify/store"
	"github.com/oussama4/commentify/web"
)

type User struct {
	store  store.Store
	logger *log.Logger
}

type UserInput struct {
	Name, Email string
}

func (u *User) routes() http.Handler {
	r := chi.NewRouter()
	r.Post("/", u.Create)
	r.Get("/{userId}", u.Get)
	r.Get("/", u.List)

	return r
}

func (u *User) List(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query()
	page, _ := web.ReadInt(qs, "page", 0)
	pageSize, _ := web.ReadInt(qs, "page_size", 0)

	users, err := u.store.ListUsers(page, pageSize)
	if err != nil {
		u.logger.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Range", "users 0-10/100")
	web.Json(w, users)
}

func (u *User) Get(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	user, err := u.store.GetUser(userId)
	if err != nil {
		if err == store.ErrNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	web.Json(w, user)
}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	user := UserInput{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	userId, err := u.store.CreateUser(user.Name, user.Email, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	web.Json(w, userId)
}