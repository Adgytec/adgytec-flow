package database

import (
	"context"
	"log"
	"os"
	"time"

	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/helpers"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

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

func createPgxConnPool() *pgxpool.Pool {
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

func (c *pgxConnection) Queries() *db_actions.Queries {
	return db_actions.New(c.connPool)
}

func (c *pgxConnection) NewTransaction(ctx context.Context) (pgx.Tx, error) {
	tx, txErr := c.connPool.Begin(ctx)
	if txErr != nil {
		return nil, txErr
	}

	actorDetails, actorDetailsErr := helpers.GetActorDetailsFromContext(ctx)
	if actorDetailsErr == nil {
		_, err := tx.Exec(ctx, `
		SELECT 
			set_config('global.actor_id', $1, true),
			set_config('global.actor_type', $2, true)`, actorDetails.ID, actorDetails.Type)
		if err != nil {
			tx.Rollback(ctx)
			return nil, err
		}
	} else {
		return nil, actorDetailsErr
	}

	return tx, nil
}

func (c *pgxConnection) Shutdown() {
	c.connPool.Close()
}

func CreatePgxDbConnectionPool() core.IDatabaseWithShutdown {
	log.Println("init db pgx pool")
	return &pgxConnection{
		connPool: createPgxConnPool(),
	}
}
