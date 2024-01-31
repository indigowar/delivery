package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func Connect(host, port, user, password, dbName string) (*pgx.Conn, error) {
	// "postgres://username:password@localhost:5432/database_name"
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbName)

	con, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return con, nil
}
