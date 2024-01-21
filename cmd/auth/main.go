package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello from AUTH service"))
	})

	_ = http.ListenAndServe(":80", nil)
}
