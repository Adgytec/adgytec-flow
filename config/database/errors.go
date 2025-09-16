package database

import "errors"

var (
	ErrRequestingTransactionInsideTransaction = errors.New("cannot start a transaction within a transaction")
)
