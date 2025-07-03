package handler

import "net/http"

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "BadRequest", http.StatusBadRequest)
}
