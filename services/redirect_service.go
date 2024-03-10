package services

import (
	"Url-Counter-Service/repositories"
	"context"
)

func CreateRedirect(code string) error {
	return repositories.CreateRedirect(code, context.Background())
}

func GetRedirects(code string) (uint, error) {
	return repositories.GetRedirectsByCode(code, context.Background())
}
