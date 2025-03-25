package scrap_paper

import (
	"context"
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
	UserId    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreatePaper(ctx context.Context, scrapPaper *ScrapPaper) error {
	row := pgsql.QueryRow(ctx, `
		INSERT INTO scrap_papers (content, is_private, user_id)
		VALUES ($1, $2, $3)
		ON CONFLICT (id) DO UPDATE SET
			content = $1,
			is_private = $2,
			user_id = $3
		RETURNING id, content, is_private, user_id, created_at, updated_at
	`, scrapPaper.Content, scrapPaper.IsPrivate, scrapPaper.UserId)

	err := row.Scan(&scrapPaper.Id, &scrapPaper.Content, &scrapPaper.IsPrivate, &scrapPaper.UserId, &scrapPaper.CreatedAt, &scrapPaper.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func GetPaper(ctx context.Context, scrapPaper *ScrapPaper) error {
	row := pgsql.QueryRow(ctx, `
		SELECT id, content, is_private, user_id, created_at, updated_at
		FROM scrap_papers
		WHERE id = $1
	`, scrapPaper.Id)

	err := row.Scan(&scrapPaper.Id, &scrapPaper.Content, &scrapPaper.IsPrivate, &scrapPaper.UserId, &scrapPaper.CreatedAt, &scrapPaper.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func UpdatePaper(ctx context.Context, scrapPaper *ScrapPaper) error {
	row := pgsql.QueryRow(ctx, `
		UPDATE scrap_papers
		SET content = $1, is_private = $2
		WHERE id = $3
		RETURNING id, content, is_private, user_id, created_at, updated_at
	`, scrapPaper.Content, scrapPaper.IsPrivate, scrapPaper.Id)

	err := row.Scan(&scrapPaper.Id, &scrapPaper.Content, &scrapPaper.IsPrivate, &scrapPaper.UserId, &scrapPaper.CreatedAt, &scrapPaper.UpdatedAt)

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
	row := pgsql.QueryRow(ctx, `
		INSERT INTO users (email, password, token)
		VALUES ($1, $2, $3)
		RETURNING id, email, password, token
	`, user.Email, user.Password, user.Token)

	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Token)
	if err != nil {
		return err
	}

	return nil
}

func GetUser(ctx context.Context, user *User) error {
	row := pgsql.QueryRow(ctx, `
		SELECT id, email, password, token
		FROM users
		WHERE id = $1
	`, user.ID)

	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Token)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(ctx context.Context, user *User) error {
	row := pgsql.QueryRow(ctx, `
		UPDATE users
		SET email = $1, password = $2, token = $3
		WHERE id = $4
		RETURNING id, email, password, token
	`, user.Email, user.Password, user.Token, user.ID)

	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Token)
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
