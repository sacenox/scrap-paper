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

type HelloWorldTemplateData struct {
	Name string
}

func HelloWorldHtml(data HelloWorldTemplateData) (string, error) {
	tmpl := template.Must(template.New("hello_world").Parse(indexRawTemplate))

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

	fmt.Println(req)

	html, err := HelloWorldHtml(HelloWorldTemplateData{Name: "World"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(html))
}
