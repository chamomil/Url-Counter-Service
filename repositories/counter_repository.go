package repositories

import (
	"Url-Counter-Service/db"
	"Url-Counter-Service/models"
	"context"
	"time"
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

func GetCounterIdByCode(code string, ctx context.Context) (uint, error) {
	var id uint
	err := db.Conn.QueryRow(ctx, `SELECT "id" FROM "counter" WHERE code = $1`, code).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func CreateRedirect(code string, ctx context.Context) error {
	id, err := GetCounterIdByCode(code, ctx)
	if err != nil {
		return err
	}
	_, err = db.Conn.Exec(ctx, `INSERT INTO "redirect" ("date", "counter_id") VALUES ($1, $2)`, time.Now(), id)
	return err
}

func GetCounters(name string, ctx context.Context) (*[]models.Counter, error) {
	var counters []models.Counter
	if name == "" {
		rows, err := db.Conn.Query(ctx, `SELECT * FROM "counter"`)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var counter models.Counter
			err := rows.Scan(&counter.Id, &counter.Url, &counter.Code, &counter.Name)
			if err != nil {
				return nil, err
			}
			counters = append(counters, counter)
		}

	} else {
		var counter models.Counter
		err := db.Conn.QueryRow(ctx, `SELECT * FROM "counter" WHERE "name" = $1`, name).Scan(&counter)
		if err != nil {
			return nil, err
		}
		counters = append(counters, counter)
	}

	return &counters, nil
}

func GetRedirectsByCode(code string, ctx context.Context) (uint, error) {
	id, err := GetCounterIdByCode(code, ctx)
	if err != nil {
		return 0, err
	}

	var count uint
	err = db.Conn.QueryRow(ctx, `SELECT COUNT(id) FROM redirect WHERE counter_id = $1 GROUP BY counter_id`, id).Scan(&count)
	return count, err
}
