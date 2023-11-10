package mariadb

import (
	"context"
	"database/sql"

	"github.com/fbriansyah/agn-aggregator-toolbox/internal/application/domain/model"
)

type DatabaseAdapter interface {
	Querier
	GetProducts(ctx context.Context, whereQuery string) ([]MProduk, error)
	GetTransaksiLog(ctx context.Context, whereQuery string) ([]model.TransaksiLog, error)
	GetTransaksiLogByID(ctx context.Context, idTrxLog int32) (model.TransaksiLog, error)
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
