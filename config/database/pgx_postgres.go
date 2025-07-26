package database

import (
	"context"
	"log"
	"os"
	"time"

	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/Adgytec/adgytec-flow/utils/interfaces"
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

func CreatePgxDbConnectionPool() interfaces.IDatabase {
	return &pgxConnection{
		connPool: createPgxConnPool(),
	}
}
