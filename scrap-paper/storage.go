package scrap_paper

import (
	"context"
	"database/sql"
	"time"
)

type ScrapPaper struct {
	Id        string    `json:"id"`
	Content   string    `json:"content"`
	IsPrivate bool      `json:"is_private"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreatePaper(ctx context.Context, scrapPaper *ScrapPaper) error {
	result, err := pgsql.Query(ctx, `
		INSERT INTO scrap_papers (content, is_private)
		VALUES ($1, $2)
		ON CONFLICT (id) DO UPDATE SET
			content = $1,
			is_private = $2
		RETURNING id, content, is_private, created_at, updated_at
	`, scrapPaper.Content, scrapPaper.IsPrivate)
	if err != nil {
		return err
	}

	if result.Next() {
		err = result.Scan(&scrapPaper.Id, &scrapPaper.Content, &scrapPaper.IsPrivate, &scrapPaper.CreatedAt, &scrapPaper.UpdatedAt)
		if err != nil {
			return err
		}
	} else {
		return sql.ErrNoRows
	}

	return nil
}

func GetPaper(ctx context.Context, scrapPaper *ScrapPaper) error {
	result, err := pgsql.Query(ctx, `
		SELECT id, content, is_private, created_at, updated_at
		FROM scrap_papers
		WHERE id = $1
	`, scrapPaper.Id)
	if err != nil {
		return err
	}

	if result.Next() {
		err = result.Scan(&scrapPaper.Id, &scrapPaper.Content, &scrapPaper.IsPrivate, &scrapPaper.CreatedAt, &scrapPaper.UpdatedAt)
		if err != nil {
			return err
		}
	} else {
		return sql.ErrNoRows
	}

	return nil
}

func UpdatePaper(ctx context.Context, scrapPaper *ScrapPaper) error {
	_, err := pgsql.Query(ctx, `
		UPDATE scrap_papers
		SET content = $1, is_private = $2
		WHERE id = $3
	`, scrapPaper.Content, scrapPaper.IsPrivate, scrapPaper.Id)

	if err != nil {
		return err
	}

	return nil
}

func DeletePaper(ctx context.Context, scrapPaper *ScrapPaper) error {
	_, err := pgsql.Exec(ctx, `
		DELETE FROM scrap_papers
		WHERE id = $1
	`, scrapPaper.Id)
	if err != nil {
		return err
	}

	return nil
}
