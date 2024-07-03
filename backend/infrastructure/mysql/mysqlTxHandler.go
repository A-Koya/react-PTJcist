package mysql

import (
	"A-Koya/react-PTJcist/adapter/repository"
	"context"
	"database/sql"
)

func (m *mysqlHandler) BeginTx(ctx context.Context) (repository.Tx, error) {
	tx, err := m.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	return newMysqlTx(tx), nil
}

type mysqlTx struct {
	tx *sql.Tx
}

func newMysqlTx(tx *sql.Tx) mysqlTx {
	return mysqlTx{tx: tx}
}

func (t mysqlTx) ExecuteContext(ctx context.Context, query string, args ...interface{}) error {
	if _, err := t.tx.ExecContext(ctx, query, args...); err != nil {
		return err
	}
	return nil
}

func (t mysqlTx) QueryContext(ctx context.Context, query string, args ...interface{}) (repository.Rows, error) {
	rows, err := t.tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	row := newMysqlRows(rows)
	return row, nil
}

func (t mysqlTx) QueryRowContext(ctx context.Context, query string, args ...interface{}) repository.Row {
	row := t.tx.QueryRowContext(ctx, query, args...)
	return newMysqlRow(row)
}

func (t mysqlTx) Commit() error {
	return t.tx.Commit()
}

func (t mysqlTx) Rollback() error {
	return t.tx.Rollback()
}
