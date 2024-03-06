package repositories

import (
	"Url-Counter-Service/db"
	"Url-Counter-Service/models"
	"Url-Counter-Service/types"
	"context"
	"fmt"
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

func GetCounters(name string, limit, offset int, ctx context.Context) (*types.PaginationResult[models.Counter], error) {
	var counters []models.Counter
	args := make([]any, 0)
	query := `SELECT * FROM "counter"`
	countQuery := `SELECT COUNT(*) AS "total" FROM "counter"`

	if name != "" {
		index := len(args) + 1

		queryStep := fmt.Sprintf(` WHERE "name" ILIKE $%d`, index)
		query += queryStep
		countQuery += queryStep

		args = append(args, name)
	}

	var total uint
	err := db.Conn.QueryRow(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		return nil, err
	}

	if limit != 0 {
		index := len(args) + 1
		query += fmt.Sprintf(` LIMIT $%d`, index)
		args = append(args, limit)
	}
	if offset != 0 {
		index := len(args) + 1
		query += fmt.Sprintf(` OFFSET $%d`, index)
		args = append(args, offset)
	}

	rows, err := db.Conn.Query(ctx, query, args...)
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

	return &types.PaginationResult[models.Counter]{
		Items: &counters, Total: total,
	}, nil
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
