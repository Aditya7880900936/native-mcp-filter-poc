package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		body, _ := io.ReadAll(r.Body)

		var payload map[string]interface{}
		json.Unmarshal(body, &payload)

		json.NewEncoder(w).Encode(map[string]any{
			"backend": "weather",
			"payload": payload,
		})
	})

	http.ListenAndServe(":8081", nil)
}