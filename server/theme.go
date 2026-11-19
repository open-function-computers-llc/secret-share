package server

import (
	"net/http"
	"os"
	"strings"
)

// theme holds the per-deployment branding. All of it comes from the env so
// nothing about the look of the site is baked into the binary.
type theme struct {
	Title       string `json:"title"`
	LogoAlt     string `json:"logoAlt"`
	ThemeColor  string `json:"themeColor"`
	Logo        string `json:"logo"` // URL of the logo, empty when there is none
	cssPath     string
	logoPath    string
	faviconPath string
}

func loadTheme() theme {
	t := theme{
		Title:       env("SITE_TITLE"),
		LogoAlt:     env("LOGO_ALT"),
		ThemeColor:  env("THEME_COLOR"),
		cssPath:     env("THEME_CSS_PATH"),
		logoPath:    env("LOGO_PATH"),
		faviconPath: env("FAVICON_PATH"),
	}
	if t.logoPath != "" {
		t.Logo = "/logo"
	}
	return t
}

func env(key string) string {
	return strings.TrimSpace(os.Getenv(key))
}

func (s *Server) handleTheme() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.sendJson(w, s.theme)
	}
}

// serveThemeFile serves an operator supplied file straight from disk, so edits
// show up without restarting the server. An unset path is a 404.
func serveThemeFile(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if path == "" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, path)
	}
}
