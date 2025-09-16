package database

import (
	"context"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/jackc/pgx/v5"
)

type pgxTx struct {
	conn    pgx.Tx
	queries *db.Queries
}

func (c *pgxTx) Queries() *db.Queries {
	return c.queries.WithTx(c.conn)
}

func (c *pgxTx) WithTransaction(_ context.Context) (Database, Tx, error) {
	return c, nil, nil
}

func newPgxTx(conn pgx.Tx, queries *db.Queries) Database {
	return &pgxTx{
		conn:    conn,
		queries: queries,
	}
}
