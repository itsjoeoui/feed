package assets

import (
	"embed"
	"net/http"

	"github.com/go-chi/chi/v5"
)

//go:embed all:dist
var Assets embed.FS

func Mount(r chi.Router) {
	r.Route("/dist", func(r chi.Router) {
		r.Handle("/*", http.FileServer(http.FS(Assets)))
	})
}
