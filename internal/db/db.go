package db

import (
	"github.com/ClickHouse/clickhouse-go/v2"
)

func New(user, pswd, addr, db string) (*clickhouse.Conn, error) {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{addr},
		Auth: clickhouse.Auth{
			Database: db,
			Username: user,
			Password: pswd,
		},
		Debug: true,
	})
	return &conn, err
}
