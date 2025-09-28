package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/utils/actor"
	"github.com/jackc/pgx/v5/pgxpool"
)

// TODO: handle this using env variables
const (
	defaultMaxConns          = int32(50)
	defaultMinConns          = int32(5)
	defaultMaxConnLifetime   = time.Hour
	defaultMaxConnIdleTime   = time.Minute * 30
	defaultHealthCheckPeriod = time.Minute
	defaultConnectTimeout    = time.Second * 5
)

func dbConfig() (*pgxpool.Config, error) {
	DATABASE_URL := os.Getenv("DB_STRING")

	dbConfig, err := pgxpool.ParseConfig(DATABASE_URL)
	if err != nil {
		return nil, ErrInvalidDBConfig
	}

	dbConfig.MaxConns = defaultMaxConns
	dbConfig.MinConns = defaultMinConns
	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

	return dbConfig, nil
}

func newPgxConnPool() (*pgxpool.Pool, error) {
	config, configErr := dbConfig()
	if configErr != nil {
		return nil, configErr
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, ErrCreatingDBConnectionPool
	}

	err = pool.Ping(context.Background())
	if err != nil {
		return nil, ErrPingingDB
	}

	return pool, nil
}

type pgxConnection struct {
	queries  *db.Queries
	connPool *pgxpool.Pool
}

func (c *pgxConnection) Queries() *db.Queries {
	return c.queries
}

func (c *pgxConnection) WithTransaction(ctx context.Context) (Database, Tx, error) {
	// Use context.Background() for transaction rollback and commit to ensure these operations complete even if the original request context is cancelled.
	tx, txErr := c.connPool.Begin(ctx)
	if txErr != nil {
		return nil, nil, txErr
	}

	// actor should be present in all the scenarios
	// sign up is created by using auth admin api so also require actor in that case too
	actorDetails, actorDetailsErr := actor.GetActorDetailsFromContext(ctx)
	if actorDetailsErr != nil {
		tx.Rollback(context.Background())
		return nil, nil, actorDetailsErr
	}

	_, err := tx.Exec(ctx, `
		SELECT 
			set_config('global.actor_id', $1, true),
			set_config('global.actor_type', $2, true)`, actorDetails.ID, actorDetails.Type)
	if err != nil {
		tx.Rollback(context.Background())
		return nil, nil, err
	}

	return newPgxTx(tx, c.Queries()), tx, nil
}

func (c *pgxConnection) Shutdown() {
	c.connPool.Close()
}

func NewPgxDbConnectionPool() (DatabaseWithShutdown, error) {
	log.Println("init db pgx pool")

	connPool, connecErr := newPgxConnPool()
	if connecErr != nil {
		return nil, connecErr
	}

	return &pgxConnection{
		connPool: connPool,
		queries:  db.New(connPool),
	}, nil
}
