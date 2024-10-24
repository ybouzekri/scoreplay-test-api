package handlers

import (
	"encoding/json"
	"net/http"
)

type responseError struct {
	code    int
	Message string `json:"message"`
}

func (e responseError) send(w http.ResponseWriter) {
	w.WriteHeader(e.code)
	json.NewEncoder(w).Encode(e)
}
