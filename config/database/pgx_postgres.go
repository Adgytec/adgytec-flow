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

func dbConfig() *pgxpool.Config {
	DATABASE_URL := os.Getenv("DB_STRING")

	dbConfig, err := pgxpool.ParseConfig(DATABASE_URL)
	if err != nil {
		log.Fatalf("error getting database config: %s", err)
	}

	dbConfig.MaxConns = defaultMaxConns
	dbConfig.MinConns = defaultMinConns
	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

	return dbConfig
}

func newPgxConnPool() *pgxpool.Pool {
	config := dbConfig()

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("error creating coneection pool: %s", err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatalf("error pinging database: %s", err)
	}

	return pool
}

type pgxConnection struct {
	connPool *pgxpool.Pool
}

func (c *pgxConnection) Queries() *db.Queries {
	return db.New(c.connPool)
}

func (c *pgxConnection) WithTransaction(ctx context.Context) (*db.Queries, Tx, error) {
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

	return c.Queries().WithTx(tx), tx, nil
}

func (c *pgxConnection) Shutdown() {
	c.connPool.Close()
}

func NewPgxDbConnectionPool() DatabaseWithShutdown {
	log.Println("init db pgx pool")
	return &pgxConnection{
		connPool: newPgxConnPool(),
	}
}
