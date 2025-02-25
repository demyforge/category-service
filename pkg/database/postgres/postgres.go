package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

const timeout = 10 * time.Second

func Connect(dsn string) (*sqlx.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return sqlx.ConnectContext(ctx, "postgres", dsn)
}

func IsDuplicate(err error) bool {
	var pqErr *pq.Error
	return errors.As(err, &pqErr) && pqErr.Code == "23505"
}
