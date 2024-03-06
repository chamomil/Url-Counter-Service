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
	return repositories.GetUrlByCode(code, context.Background())
}

func GetCounters(name string) (*[]models.Counter, error) {
	return repositories.GetCounters(name, context.Background())
}
