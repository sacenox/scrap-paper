package scrap_paper_test

import (
	"context"
	"testing"

	scrap_paper "encore.app/scrap-paper"
)

func TestScrapPaperCRUD(t *testing.T) {
	ctx := context.Background()

	scrapPaper := &scrap_paper.ScrapPaper{
		Content: "Hello test, test.",
		IsPrivate: false,
	}

	t.Run("Create", func(t *testing.T) {
		err := scrap_paper.Create(ctx, scrapPaper)
		if err != nil {
			t.Fatalf("failed to create scrap paper: %v", err)
		}

		if scrapPaper.Id == "" {
			t.Fatalf("scrap paper id is empty")
		}
	
		if scrapPaper.Content != "Hello test, test." {
			t.Fatalf("scrap paper content is not Hello test, test.: %v", scrapPaper.Content)
		}
	
		if scrapPaper.IsPrivate != false {
			t.Fatalf("scrap paper private is not false: %v", scrapPaper.IsPrivate)
		}
	})

	t.Run("Get", func(t *testing.T) {
		err := scrap_paper.Get(ctx, scrapPaper)
		if err != nil {
			t.Fatalf("failed to get scrap paper: %v", err)
		}

		if scrapPaper.Id == "" {
			t.Fatalf("scrap paper id is empty")
		}
		
		if scrapPaper.Content != "Hello test, test." {
			t.Fatalf("scrap paper content is not Hello test, test.: %v", scrapPaper.Content)
		}

		if scrapPaper.IsPrivate != false {
			t.Fatalf("scrap paper private is not false: %v", scrapPaper.IsPrivate)
		}
	})
}
