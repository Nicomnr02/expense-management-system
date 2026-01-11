package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Tx pgx.Tx

type Transaction interface {
	Begin(c context.Context) (Tx, error)
	Commit(c context.Context, tx Tx) error
	Rollback(c context.Context, tx Tx) error
}

type transaction struct {
	*Database
}

func NewTransaction(conn *Database) Transaction {
	return &transaction{conn}
}

func (t *transaction) Begin(c context.Context) (Tx, error) {
	return t.Conn.Begin(c)
}

func (t *transaction) Commit(c context.Context, tx Tx) error {
	return tx.Commit(c)
}

func (t *transaction) Rollback(c context.Context, tx Tx) error {
	return tx.Rollback(c)
}
