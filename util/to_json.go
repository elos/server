package util

import (
	"encoding/json"
	"net/http"
)

func ToJSON(v interface{}) ([]byte, error) {
	// Always pretty-print JSON
	return json.MarshalIndent(v, "", "    ")
}

func ContentJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}
