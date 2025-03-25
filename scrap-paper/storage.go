package scrap_paper

import (
	"context"
	"database/sql"
	"time"
)

type User struct {
	ID       string
	Email    string
	Token    string
	Password string
}

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

func CreateUser(ctx context.Context, user *User) error {
	result, err := pgsql.Query(ctx, `
		INSERT INTO users (email, password)
		VALUES ($1, $2)
		RETURNING id, email, password
	`, user.Email, user.Password)
	if err != nil {
		return err
	}

	if result.Next() {
		err = result.Scan(&user.ID, &user.Email, &user.Password)
		if err != nil {
			return err
		}
	} else {
		return sql.ErrNoRows
	}

	return nil
}

func GetUser(ctx context.Context, user *User) error {
	result, err := pgsql.Query(ctx, `
		SELECT id, email, password, token
		FROM users
		WHERE id = $1
	`, user.ID)
	if err != nil {
		return err
	}

	if result.Next() {
		err = result.Scan(&user.ID, &user.Email, &user.Password, &user.Token)
		if err != nil {
			return err
		}
	} else {
		return sql.ErrNoRows
	}

	return nil
}

func UpdateUser(ctx context.Context, user *User) error {
	_, err := pgsql.Exec(ctx, `
		UPDATE users
		SET email = $1, password = $2, token = $3
		WHERE id = $4
	`, user.Email, user.Password, user.Token, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(ctx context.Context, user *User) error {
	_, err := pgsql.Exec(ctx, `
		DELETE FROM users
		WHERE id = $1
	`, user.ID)
	if err != nil {
		return err
	}

	return nil
}
