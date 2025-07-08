package handler

import "net/http"

type ErrorHandler struct{}

func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{}
}

func (e *ErrorHandler) BadRequest(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "BadRequest", http.StatusBadRequest)
}
