package database

import (
	"context"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/jackc/pgx/v5"
)

type txStub struct{}

func (tx txStub) Commit(_ context.Context) error {
	return nil
}

func (tx txStub) Rollback(_ context.Context) error {
	return nil
}

type pgxTx struct {
	queries *db.Queries
}

func (c *pgxTx) Queries() *db.Queries {
	return c.queries
}

// WithTransaction() return parent transaction with stub for Tx interface
// actual transaction methods will be handled by parent action
func (c *pgxTx) WithTransaction(_ context.Context) (Database, Tx, error) {
	return c, txStub{}, nil
}

func newPgxTx(conn pgx.Tx, queries *db.Queries) Database {
	return &pgxTx{
		queries: queries.WithTx(conn),
	}
}
