package mariadb

import (
	"context"
	"database/sql"
)

type DatabaseAdapter interface {
	Querier
	GetProducts(ctx context.Context, whereQuery string) ([]MProduk, error)
}

type DatabaseStore struct {
	db *sql.DB
	*Queries
}

func NewDatabaseAdapter(db *sql.DB) DatabaseAdapter {
	return &DatabaseStore{
		db:      db,
		Queries: New(db),
	}
}
