package web

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Json encodes v to JSON and writes it to the client
func Json(w http.ResponseWriter, statusCode int, v map[string]interface{}) error {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	if err := enc.Encode(v); err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	if _, err := w.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}
