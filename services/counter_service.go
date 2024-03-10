package services

import (
	"Url-Counter-Service/models"
	"Url-Counter-Service/repositories"
	"Url-Counter-Service/types"
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

func GetCounters(name string, limit int, offset int) (*types.PaginationResult[models.Counter], error) {
	return repositories.GetCounters(name, limit, offset, context.Background())
}
