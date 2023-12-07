package repository

import (
	"context"
	"database/sql"
	"roomate/config"
)

type Database interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func NewRepository(db config.InfraConfig) *Queries {
	return &Queries{db: db}
}

type Queries struct {
	db Database
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db: tx,
	}
}
