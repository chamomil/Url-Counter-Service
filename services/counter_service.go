package services

import (
	"Url-Counter-Service/models"
	"Url-Counter-Service/repositories"
	"context"
	"github.com/google/uuid"
)

func CreateCounter(counter *models.Counter) error {
	counter.Code = uuid.New().String()
	return repositories.CreateCounter(counter, context.Background())
}

func GetUrlByCode(code string) (string, error) {
	url, err := repositories.GetUrlByCode(code, context.Background())
	if err != nil {
		return "", err
	}
	err = CreateRedirect(code)
	return url, err
}

func GetCounters(name string) (*[]models.Counter, error) {
	return repositories.GetCounters(name, context.Background())
}

func CreateRedirect(code string) error {
	return repositories.CreateRedirect(code, context.Background())
}

func GetRedirects(code string) (uint, error) {
	return repositories.GetRedirectsByCode(code, context.Background())
}
