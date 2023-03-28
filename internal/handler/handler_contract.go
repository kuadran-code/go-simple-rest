package handler

import (
	"encoding/json"
	"net/http"
)

type Handler struct{}

type dataError struct {
	Message string `json:"message"`
}

func (h *Handler) ResponseError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	d, _ := json.Marshal(dataError{
		Message: message,
	})
	w.Write(d)
}

func (h *Handler) Response(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	d, _ := json.Marshal(data)
	w.Write(d)
}
