package handlers

import (
	"encoding/json"
	"net/http"
)

func HelloAPI() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		helloServer := "Welcome to my blog api"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(helloServer)
	}
}
