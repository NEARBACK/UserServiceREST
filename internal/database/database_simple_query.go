package database

import (
	"context"
	"useservice/internal/definitions"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Database is IDatabase implementation
type Database struct {
	readPool  *pgxpool.Pool
	writePool *pgxpool.Pool
	log       definitions.Logger
}

func (dbl *Database) GetContext() context.Context { return context.Background() }

func (dbl *Database) Get(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error {
	// without cache
	return pgxscan.Select(ctx, dbl.readPool, returnValue, sql, args...)
}

func (dbl *Database) GetOne(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error {
	return pgxscan.Get(ctx, dbl.readPool, returnValue, sql, args...)
}

func (dbl *Database) Count(ctx context.Context, returnValue *int64, sql string, args ...interface{}) error {
	err := dbl.readPool.QueryRow(ctx, sql, args...).Scan(&returnValue)
	if err != nil {
		return err
	}

	return nil
}

func (dbl *Database) Insert(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error {
	conn, err := dbl.writePool.Acquire(ctx)
	defer func() {
		if conn != nil {
			conn.Release()
		}
	}()
	if err != nil {
		return err
	}

	if returnValue != nil {
		if err := pgxscan.Get(ctx, conn, returnValue, sql, args...); err != nil {
			return err
		}
	} else {
		cmd, err := conn.Exec(ctx, sql, args...)
		if err != nil {
			return err
		}
		if !cmd.Insert() {
			dbl.log.Error(err)
		}
	}

	return nil
}

func (dbl *Database) Update(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error {
	conn, err := dbl.writePool.Acquire(ctx)
	defer func() {
		if conn != nil {
			conn.Release()
		}
	}()
	if err != nil {
		return err
	}

	if returnValue != nil {
		if err := pgxscan.Get(ctx, conn, returnValue, sql, args...); err != nil {
			return err
		}
	} else {
		cmd, err := conn.Exec(ctx, sql, args...)
		if err != nil {
			return err
		}
		if !cmd.Update() {
			dbl.log.Error(err)
		}
	}

	return nil
}

func (dbl *Database) Delete(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error {
	conn, err := dbl.writePool.Acquire(ctx)
	defer func() {
		if conn != nil {
			conn.Release()
		}
	}()
	if err != nil {
		return err
	}

	if returnValue != nil {
		if err := pgxscan.Get(ctx, conn, returnValue, sql, args...); err != nil {
			return err
		}
	} else {
		cmd, err := conn.Exec(ctx, sql, args...)
		if err != nil {
			return err
		}
		if !cmd.Delete() {
			dbl.log.Error(err)
		}
	}

	return nil
}

func (dbl *Database) CloseAll() {
	dbl.writePool.Close()
	dbl.readPool.Close()
}

func (dbl *Database) PingAll() (err error) {

	err = dbl.writePool.Ping(context.Background())
	if err != nil {
		return err
	}
	err = dbl.readPool.Ping(context.Background())
	if err != nil {
		return err
	}

	return err
}

func (dbl *Database) StartTransaction(ctx context.Context, txLevel pgx.TxIsoLevel) (tx ITransaction, err error) {

	t, err := dbl.writePool.BeginTx(ctx, pgx.TxOptions{IsoLevel: txLevel})
	if err != nil {
		return nil, err
	}

	return &Transaction{t: t}, nil
}
