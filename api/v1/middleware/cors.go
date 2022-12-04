package middleware

import (
	"net/http"

	"github.com/burntcarrot/hashsearch/pkg/config"
)

// CORS middleware.
func CORS(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Header().Set("Access-Control-Allow-Origin", config.CORS_ALLOW_ORIGIN)
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		return
	}

	next(w, r)
}
