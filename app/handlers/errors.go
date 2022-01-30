package handlers

import (
	"log"
	"net/http"

	"github.com/oussama4/commentify/base/web"
)

func respondError(l *log.Logger, w http.ResponseWriter, statusCode int, msg string) {
	l.Println("ERROR: ", msg)
	if err := web.Json(w, statusCode, map[string]interface{}{"error": msg}); err != nil {
		l.Println("ERROR: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}
