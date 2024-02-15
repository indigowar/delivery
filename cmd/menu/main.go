package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("MENU SVC"))
	})

	_ = http.ListenAndServe(":80", nil)
}