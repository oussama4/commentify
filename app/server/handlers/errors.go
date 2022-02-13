package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/oussama4/commentify/base/validate"
	"github.com/oussama4/commentify/base/web"
)

func respondError(l *log.Logger, w http.ResponseWriter, statusCode int, err error) {
	l.Println("ERROR: ", err)

	var validationErr validate.ValidationError
	if errors.As(err, &validationErr) {
		if err := web.Json(w, statusCode, map[string]interface{}{"errors": validationErr.Fields()}); err != nil {
			l.Println("ERROR: ", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	if err := web.Json(w, statusCode, map[string]interface{}{"errors": err.Error()}); err != nil {
		l.Println("ERROR: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}
