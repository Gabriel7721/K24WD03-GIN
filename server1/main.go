package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// Set header = JSON
		w.Header().Set("content-type", "application/json")
		response := map[string]string{"message": "pong"}
		json.NewEncoder(w).Encode(response)
	})
	http.ListenAndServe(":9999", nil)
}
