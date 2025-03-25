package scrap_paper_test

import (
	"context"
	"database/sql"
	"testing"

	scrap_paper "encore.app/scrap-paper"
	"encore.dev/et"
)

func TestScrapPaperCRUD(t *testing.T) {
	ctx := context.Background()

	if _, err := et.NewTestDatabase(ctx, "scrap_paper"); err != nil {
		t.Fatalf("failed to create test database: %v", err)
	}

	scrapPaper := &scrap_paper.ScrapPaper{
		Content:   "Hello test, test.",
		IsPrivate: false,
	}

	t.Run("Create", func(t *testing.T) {
		err := scrap_paper.CreatePaper(ctx, scrapPaper)
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
		err := scrap_paper.GetPaper(ctx, scrapPaper)
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

	t.Run("Update", func(t *testing.T) {
		scrapPaper.Content = "Hello test, test. updated."
		err := scrap_paper.UpdatePaper(ctx, scrapPaper)
		if err != nil {
			t.Fatalf("failed to update scrap paper: %v", err)
		}

		if scrapPaper.Content != "Hello test, test. updated." {
			t.Fatalf("scrap paper content is not Hello test, test. updated.: %v", scrapPaper.Content)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		err := scrap_paper.DeletePaper(ctx, scrapPaper)
		if err != nil {
			t.Fatalf("failed to delete scrap paper: %v", err)
		}

		deleted := scrap_paper.ScrapPaper{}
		err = scrap_paper.GetPaper(ctx, &deleted)

		if err != sql.ErrNoRows {
			t.Fatalf("Row should not be found: %v", err)
		}
	})
}
