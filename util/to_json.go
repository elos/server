package util

import (
	"encoding/json"
	"net/http"
)

func ToJson(v interface{}) ([]byte, error) {
	// Always pretty-print JSON
	return json.MarshalIndent(v, "", "    ")
}

func ContentJson(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}
