package scrap_paper_test

import (
	"context"
	"testing"

	scrap_paper "encore.app/scrap-paper"
	"encore.dev/et"
)

func TestUserCRUD(t *testing.T) {
	ctx := context.Background()

	if _, err := et.NewTestDatabase(ctx, "scrap_paper"); err != nil {
		t.Fatalf("failed to create test database: %v", err)
	}

	user := &scrap_paper.User{
		Email:    "test@test.com",
		Password: "password",
		Token:    "",
	}

	t.Run("Create", func(t *testing.T) {
		err := scrap_paper.CreateUser(ctx, user)
		if err != nil {
			t.Fatalf("failed to create user: %v", err)
		}

		if user.ID == "" {
			t.Fatalf("user id is empty")
		}

		if user.Email != "test@test.com" {
			t.Fatalf("user email is not test@test.com: %v", user.Email)
		}

		if user.Password != "password" {
			t.Fatalf("user password is not password: %v", user.Password)
		}

		if user.Token != "" {
			t.Fatalf("user token is not empty: %v", user.Token)
		}
	})

	t.Run("Get", func(t *testing.T) {
		retrieved := &scrap_paper.User{}
		retrieved.ID = user.ID

		err := scrap_paper.GetUser(ctx, retrieved)
		if err != nil {
			t.Fatalf("failed to get user: %v", err)
		}

		if retrieved.Email != "test@test.com" {
			t.Fatalf("retrieved user email is not test@test.com: %v", retrieved.Email)
		}

		if retrieved.Password != "password" {
			t.Fatalf("retrieved user password is not password: %v", retrieved.Password)
		}
	})

	t.Run("Update", func(t *testing.T) {
		user.Email = "updated@test.com"
		user.Password = "newpassword"
		user.Token = "sometoken"

		err := scrap_paper.UpdateUser(ctx, user)
		if err != nil {
			t.Fatalf("failed to update user: %v", err)
		}

		retrieved := &scrap_paper.User{}
		retrieved.ID = user.ID

		err = scrap_paper.GetUser(ctx, retrieved)
		if err != nil {
			t.Fatalf("failed to get updated user: %v", err)
		}

		if retrieved.Email != "updated@test.com" {
			t.Fatalf("updated user email is not updated@test.com: %v", retrieved.Email)
		}

		if retrieved.Password != "newpassword" {
			t.Fatalf("updated user password is not newpassword: %v", retrieved.Password)
		}

		if retrieved.Token != "sometoken" {
			t.Fatalf("updated user token is not sometoken: %v", retrieved.Token)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		err := scrap_paper.DeleteUser(ctx, user)
		if err != nil {
			t.Fatalf("failed to delete user: %v", err)
		}

		deleted := &scrap_paper.User{}
		deleted.ID = user.ID

		err = scrap_paper.GetUser(ctx, deleted)

		if err == nil {
			t.Fatalf("Row should not be found: %v", err)
		}
	})
}

func TestScrapPaperCRUD(t *testing.T) {
	ctx := context.Background()

	if _, err := et.NewTestDatabase(ctx, "scrap_paper"); err != nil {
		t.Fatalf("failed to create test database: %v", err)
	}

	user := &scrap_paper.User{
		Email:    "test@test.com",
		Password: "password",
		Token:    "",
	}
	err := scrap_paper.CreateUser(ctx, user)
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	scrapPaper := &scrap_paper.ScrapPaper{
		Content:   "Hello test, test.",
		IsPrivate: false,
		UserId:    user.ID,
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

		if err == nil {
			t.Fatalf("Row should not be found: %v", err)
		}
	})
}
