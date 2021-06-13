package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oussama4/commentify/store"
)

type User struct {
	store store.Store
}

type UserInput struct {
	Name, Email string
}

func (u *User) routes() http.Handler {
	r := chi.NewRouter()
	r.Post("/", u.Create)

	return r
}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	user := UserInput{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	userId, err := u.store.CreateUser(user.Name, user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	if err := enc.Encode(userId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(buf.Bytes())
}
