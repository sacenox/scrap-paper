package scrap_paper

import (
	"net/http"

	"encore.dev"
)

//encore:api public raw path=/assets/:filename
func Assets(w http.ResponseWriter, r *http.Request) {
	// This path is relative to the root of the project
	path := "." + encore.CurrentRequest().Path
	http.ServeFile(w, r, path)
}
