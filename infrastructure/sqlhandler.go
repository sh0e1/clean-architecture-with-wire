package infrastructure

import (
	"context"
	"database/sql"

	"github.com/sh0e1/wire/interface/database"
)

type Driver string

const (
	SQLite Driver = "sqlite3"
)

func NewSQLHandler(driver Driver, datasource string) (*SQLHandler, error) {
	db, err := sql.Open(string(SQLite), datasource)
	if err != nil {
		return nil, err
	}
	return &SQLHandler{db}, nil
}

type SQLHandler struct {
	db *sql.DB
}

func (h *SQLHandler) ExecuteContext(ctx context.Context, query string, args ...interface{}) (database.Result, error) {
	result, err := h.db.ExecContext(ctx, query, args...)
	return result, err
}

func (h *SQLHandler) QueryContext(ctx context.Context, query string, args ...interface{}) (database.Rows, error) {
	rows, err := h.db.QueryContext(ctx, query, args...)
	return rows, err
}

func (h *SQLHandler) QueryRowContext(ctx context.Context, query string, args ...interface{}) database.Row {
	row := h.db.QueryRowContext(ctx, query, args...)
	return row
}

func (h *SQLHandler) Close() {
	_ = h.db.Close()
}
