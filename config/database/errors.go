package database

import "errors"

var (
	ErrRequestingTransactionInsideTransaction = errors.New("error transaction inside transaction")
)
