package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		backend := r.Header.Get("x-tool-backend")

		fmt.Println("AUTH CHECK =>", backend)

		if backend == "github" {

			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("github tools blocked"))

			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("allowed"))
	})

	http.ListenAndServe(":9000", nil)
}
