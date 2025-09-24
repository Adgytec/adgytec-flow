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

// WithTransaction returns the parent transaction database connection and a no-op Tx stub.
// This is to prevent nested transactions, as the actual commit/rollback will be handled by the parent transaction.
func (c *pgxTx) WithTransaction(_ context.Context) (Database, Tx, error) {
	return c, txStub{}, nil
}

func newPgxTx(conn pgx.Tx, queries *db.Queries) Database {
	return &pgxTx{
		queries: queries.WithTx(conn),
	}
}
