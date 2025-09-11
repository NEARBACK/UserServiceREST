package database

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
)

// Transaction wraps pgx.Tx
type Transaction struct {
	t pgx.Tx
}

func (tx *Transaction) TxGetContext() context.Context { return context.Background() }

func (tx *Transaction) TxGet(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error {
	rows, err := tx.t.Query(ctx, sql, args...)
	if err != nil {
		return err
	}
	return pgxscan.ScanAll(returnValue, rows)
}

func (tx *Transaction) TxGetOne(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error {
	rows, err := tx.t.Query(ctx, sql, args...)
	if err != nil {
		return err
	}
	return pgxscan.ScanOne(returnValue, rows)
}

func (tx *Transaction) TxCount(ctx context.Context, returnValue *int64, sql string, args ...interface{}) error {
	return tx.t.QueryRow(ctx, sql, args...).Scan(&returnValue)
}

func (tx *Transaction) TxInsert(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error {
	if returnValue != nil {

		rows, err := tx.t.Query(ctx, sql, args...)
		if err != nil {
			return err
		}
		return pgxscan.ScanOne(returnValue, rows)
	} else {
		_, err := tx.t.Exec(ctx, sql, args...)
		return err
	}
}

func (tx *Transaction) TxUpdate(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error {
	if returnValue != nil {
		rows, err := tx.t.Query(ctx, sql, args...)
		if err != nil {
			return err
		}
		return pgxscan.ScanAll(returnValue, rows)
	} else {
		_, err := tx.t.Exec(ctx, sql, args...)
		return err
	}
}

func (tx *Transaction) TxDelete(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error {
	if returnValue != nil {
		return tx.t.QueryRow(ctx, sql, args...).Scan(&returnValue)
	} else {
		_, err := tx.t.Exec(ctx, sql, args...)
		return err
	}
}

func (tx *Transaction) TxCommit(ctx context.Context) (err error) {
	return tx.t.Commit(ctx)
}

func (tx *Transaction) TxRollback(ctx context.Context) (err error) {
	return tx.t.Rollback(ctx)
}
