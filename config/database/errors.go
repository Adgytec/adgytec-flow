package database

import "errors"

var (
	ErrInvalidDBConfig          = errors.New("invalid db config")
	ErrCreatingDBConnectionPool = errors.New("can't create db connection pool")
	ErrPingingDB                = errors.New("can't ping db")
)
