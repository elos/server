package util

import "net/http"

func ResourceResponse(w http.ResponseWriter, status int, resource interface{}) {
	ContentJson(w)

	w.WriteHeader(status)

	bytes, _ := ToJson(resource)

	w.Write(bytes)
}
