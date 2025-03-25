package scrap_paper_test

import (
	"strings"
	"testing"

	"encore.app/lib"
	scrap_paper "encore.app/scrap-paper"
)

// request and check the html for the htmx script tag and then for the tailwind css link
// and finally for the title.
func TestIndex(t *testing.T) {
	// We can't use requests to test a raw endpoint so we just check the html
	body, err := lib.RenderTemplate(scrap_paper.IndexRawTemplate, scrap_paper.IndexTemplateData{
		Title: "Scrap Paper",
		Env:   "test",
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Run("htmx script tag", func(t *testing.T) {
		if !strings.Contains(string(body), "htmx.org") {
			t.Fatal("htmx script tag not found")
		}
	})

	t.Run("tailwind css link", func(t *testing.T) {
		if !strings.Contains(string(body), "@tailwindcss") {
			t.Fatal("tailwind css link not found")
		}
	})

	t.Run("title", func(t *testing.T) {
		if !strings.Contains(string(body), "Scrap Paper") {
			t.Fatal("title not found")
		}
	})
}
