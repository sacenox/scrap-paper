package scrap_paper

import (
	_ "embed"
	"fmt"
	"net/http"

	"encore.app/lib"
	"encore.dev"
)

//go:embed index.html
var IndexRawTemplate string

type IndexTemplateData struct {
	Title string
	Env   string
}

//encore:api public raw path=/!fallback
func (svc *ScrapPaperService) Index(w http.ResponseWriter, r *http.Request) {
	req := encore.CurrentRequest()
	env := encore.Meta().Environment.Name
	title := "Scrap Paper - Home"

	if env != "prod" {
		title = fmt.Sprintf("%s [%s]", title, env)
	}

	if req.Path == "/" {
		html, err := lib.RenderTemplate(IndexRawTemplate, IndexTemplateData{
			Title: title,
			Env:   env,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return // TODO: Return an error page
		}
		w.Write([]byte(html))
		return // DONE
	}

	// Since we are abusing fallback handler to mount a page on "/"
	// we need to return a 404 error if the path is not "/"
	http.NotFound(w, r)
}
