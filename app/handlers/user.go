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

type User struct {
	store  store.Store
	logger *log.Logger
}

type UserInput struct {
	Name, Email string
}

func (ui *UserInput) valid() error {
	v := validate.New()

	v.Check(ui.Name != "", "title", "user name is required")
	v.Check(validate.EmailFormat(ui.Email), "email", "invalid email")

	return v.Valid()
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
	pageSize, _ := web.ReadInt(qs, "page_size", 10)

	users, err := u.store.ListUsers(page, pageSize)
	if err != nil {
		respondError(u.logger, w, http.StatusInternalServerError, err)
	}
	web.Json(w, http.StatusOK, map[string]interface{}{"users": users})
}

func (u *User) Get(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	user, err := u.store.GetUser(userId)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			respondError(u.logger, w, http.StatusNotFound, err)
		}
		respondError(u.logger, w, http.StatusInternalServerError, err)
	}

	web.Json(w, http.StatusOK, map[string]interface{}{"user": user})
}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	user := UserInput{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondError(u.logger, w, http.StatusInternalServerError, err)
	}

	if err := user.valid(); err != nil {
		respondError(u.logger, w, http.StatusUnprocessableEntity, err)
	}

	userId, err := u.store.CreateUser(user.Name, user.Email)
	if err != nil {
		respondError(u.logger, w, http.StatusInternalServerError, err)
	}

	web.Json(w, http.StatusOK, map[string]interface{}{"user_id": userId})
}
