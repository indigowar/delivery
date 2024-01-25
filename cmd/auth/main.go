package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("AUTH SERVICE"))
	})

	_ = http.ListenAndServe(":80", nil)
}

