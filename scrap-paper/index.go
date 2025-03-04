package scrap_paper

// We are building a service that uses htmx to return html from a url
import (
	"bytes"
	_ "embed"
	"fmt"
	"net/http"

	"html/template"

	"encore.dev"
)

//go:embed index.html
var indexRawTemplate string

type IndexTemplateData struct {
	Title string
	Env   string
}

func IndexHtml(data IndexTemplateData) (string, error) {
	tmpl := template.Must(template.New("index").Parse(indexRawTemplate))

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

//
// Warning: This is a fallback handler, being missused as a catch-all handler.
// It is used to handle requests that don't match any other route.
// In our case it means it renders the index page.
//

//encore:api public raw path=/!fallback
func Index(w http.ResponseWriter, r *http.Request) {
	req := encore.CurrentRequest()
	env := encore.Meta().Environment.Name
	title := "Scrap Paper - Home"

	if env != "prod" {
		title = fmt.Sprintf("%s [%s]", title, env)
	}

	if req.Path == "/" {
		html, err := IndexHtml(IndexTemplateData{
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
	// TODO: Return an error page
	http.NotFound(w, r)
}

//
// Handle static assets
//

//encore:api public raw path=/assets/:filename
func Assets(w http.ResponseWriter, r *http.Request) {
	// This path is relative to the root of the project
	path := "." + encore.CurrentRequest().Path
	http.ServeFile(w, r, path)
}
