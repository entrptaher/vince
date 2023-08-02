package assets

import (
	"embed"
	"io/fs"
	"net/http"
	"strings"

	"github.com/vinceanalytics/vince/internal/must"
	"github.com/vinceanalytics/vince/internal/plug"
)

var files = map[string]bool{
	"/favicon.svg": true,
	"/favicon.ico": true,
	"/favicon":     true,
	"/robots.txt":  true,
	"/logo.svg":    true,
	"/index.html":  true,
	"/":            true,
}

//go:embed ui
var static embed.FS

var ui = must.Must(fs.Sub(static, "ui"))()

func match(path string) bool {
	return strings.HasPrefix(path, "/static") ||
		files[path]
}

func Plug() plug.Plug {
	app := http.FileServer(http.FS(ui))
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if match(r.URL.Path) {
				app.ServeHTTP(w, r)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}
