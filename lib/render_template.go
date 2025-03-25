package lib

import (
	"bytes"
	"html/template"
)

func RenderTemplate(tmpl string, data interface{}) (string, error) {
	t := template.Must(template.New("index").Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
