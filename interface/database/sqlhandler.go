package database

import "context"

type SQLHandler interface {
	ExecuteContext(ctx context.Context, query string, args ...interface{}) (Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) Row
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Rows interface {
	Next() bool
	Err() error
	Scan(dest ...interface{}) error
	Close() error
}

type Row interface {
	Scan(dest ...interface{}) error
	Err() error
}
