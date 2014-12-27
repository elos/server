package util

import "net/http"

func WriteResourceResponse(w http.ResponseWriter, status int, resource interface{}) {
	w.WriteHeader(status)
	WriteJSON(w, resource)
}
