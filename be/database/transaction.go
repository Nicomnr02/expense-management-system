package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Tx pgx.Tx

type Transaction struct {
	*Database
}

func NewTransaction(conn *Database) *Transaction {
	return &Transaction{conn}
}

func (t *Transaction) Begin(c context.Context) (Tx, error) {
	return t.Conn.Begin(c)
}

func (t *Transaction) Commit(c context.Context, tx Tx) error {
	return tx.Commit(c)
}

func (t *Transaction) Rollback(c context.Context, tx Tx) error {
	return tx.Rollback(c)
}
