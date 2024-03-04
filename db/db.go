package db

import (
	"Url-Counter-Service/config"
	"context"
	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

func InitConnection(c context.Context, config *config.PostgresConfig) error {
	connConfig, _ := pgx.ParseConfig("")
	connConfig.Host = config.Host
	connConfig.Port = config.Port
	connConfig.User = config.User
	connConfig.Password = config.Password
	connConfig.Database = config.Database

	var err error
	conn, err = pgx.ConnectConfig(c, connConfig)
	if err != nil {
		return err
	}
	return nil
}

func CloseConnection(c context.Context) error {
	return conn.Close(c)
}
