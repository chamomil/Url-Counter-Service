package repositories

import (
	"Url-Counter-Service/db"
	"context"
	"time"
)

func CreateRedirect(code string, ctx context.Context) error {
	var id uint
	err := db.Conn.QueryRow(ctx, `SELECT "id" FROM "counter" WHERE code = $1`, code).Scan(&id)
	if err != nil {
		return err
	}
	_, err = db.Conn.Exec(ctx, `INSERT INTO "redirect" ("date", "counter_id") VALUES ($1, $2)`, time.Now(), id)
	return err
}

func GetRedirectsByCode(code string, ctx context.Context) (uint, error) {
	var count uint
	err := db.Conn.QueryRow(ctx,
		`SELECT COUNT(id) FROM redirect JOIN counter ON redirect.counter_id = counter.id WHERE counter.code = $1 GROUP BY counter_id`,
		code).Scan(&count)
	return count, err
}
