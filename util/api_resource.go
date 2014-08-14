package util

import "net/http"

func ResourceResponse(w http.ResponseWriter, status int, resource interface{}) {
	ContentJSON(w)

	w.WriteHeader(status)

	bytes, _ := ToJSON(resource)

	w.Write(bytes)
}
