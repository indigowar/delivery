package app

import (
	"net/http"

	"github.com/indigowar/delivery/internal/accounts/config"
)

func Run(_ *config.Config) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello, from ACCOUNTS service"))
	})

	_ = http.ListenAndServe(":80", nil)
}
