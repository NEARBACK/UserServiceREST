package database

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type IDatabase interface {
	Get(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error
	GetOne(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error
	Count(ctx context.Context, returnValue *int64, sql string, args ...interface{}) error
	Insert(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error
	Update(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error
	Delete(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error
	GetContext() context.Context
	CloseAll()
	PingAll() error

	StartTransaction(ctx context.Context, isoLevel pgx.TxIsoLevel) (tx ITransaction, err error)
}

type ITransaction interface {
	TxGet(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error
	TxGetOne(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error
	TxCount(ctx context.Context, returnValue *int64, sql string, args ...interface{}) error
	TxInsert(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error
	TxUpdate(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error
	TxDelete(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error
	TxGetContext() context.Context

	TxCommit(ctx context.Context) (err error)
	TxRollback(ctx context.Context) (err error)
}
