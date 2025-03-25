package scrap_paper

import (
	"context"

	"encore.dev/beta/auth"
)

//encore:authhandler
func AuthHandler(ctx context.Context, token string) (auth.UID, *User, error) {
	// TODO: Check if the token is set on the db

	return "123", &User{
		ID:    "123",
		Email: "test@test.com",
		Token: token,
	}, nil
}

type LoginParams struct {
	Email    string
	Password string
}

type LoginResponse struct {
	Token string
}

//encore:api public method=POST path=/auth/login
func Login(ctx context.Context, params LoginParams) (LoginResponse, error) {
	// TODO: Implement the login, return the token
	return LoginResponse{
		Token: "success",
	}, nil
}
