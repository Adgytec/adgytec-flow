package database

import (
	"context"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/jackc/pgx/v5"
)

type pgxTx struct {
	queries *db.Queries
}

func (c *pgxTx) Queries() *db.Queries {
	return c.queries
}

// WithTransaction prevents nested transactions by always returning ErrRequestingTransactionInsideTransaction.
// Callers can check for this error to know they are already within a transaction.
func (c *pgxTx) WithTransaction(_ context.Context) (Database, Tx, error) {
	return nil, nil, ErrRequestingTransactionInsideTransaction
}

func newPgxTx(conn pgx.Tx, queries *db.Queries) Database {
	return &pgxTx{
		queries: queries.WithTx(conn),
	}
}
