package repository

import "database/sql"

//go:generate mockgen -source=$GOFILE -destination=./mocks/mock_$GOFILE
type TransactionExecutor interface {
	WithTransaction(fn func(tx *sql.Tx) error) error
}
