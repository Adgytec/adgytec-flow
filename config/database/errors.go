package database

import (
	"fmt"
)

type InvalidDBConfigError struct {
	cause error
}

func (e *InvalidDBConfigError) Error() string {
	return fmt.Sprintf("error parsing db config: %s", e.cause)
}

type CreatingDBConnectionPoolError struct {
	cause error
}

func (e *CreatingDBConnectionPoolError) Error() string {
	return fmt.Sprintf("error creating db connection pool: %s", e.cause)
}

type PingingDBError struct {
	cause error
}

func (e *PingingDBError) Error() string {
	return fmt.Sprintf("error pinging db: %s", e.cause)
}
