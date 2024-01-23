package app

import (
	"net/http"

	"github.com/indigowar/delivery/internal/auth/config"
)

func Run(cfg *config.Config) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello from AUTH service"))
	})

	_ = http.ListenAndServe(":80", nil)
}
