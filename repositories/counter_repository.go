package repositories

import (
	"Url-Counter-Service/db"
	"Url-Counter-Service/models"
	"context"
)

func CreateCounter(counter *models.Counter, ctx context.Context) error {
	var id uint

	err := db.Conn.QueryRow(ctx, `INSERT INTO "counter" ("url", "code", "name") VALUES ($1, $2, $3) RETURNING "id"`, counter.Url, counter.Code, counter.Name).Scan(&id)
	if err != nil {
		return err
	}
	counter.Id = id

	return nil
}

func GetUrlByCode(code string, ctx context.Context) (string, error) {
	var url string
	err := db.Conn.QueryRow(ctx, `SELECT "url" FROM "counter" WHERE code = $1`, code).Scan(&url)
	if err != nil {
		return "", err
	}
	return url, nil
}
