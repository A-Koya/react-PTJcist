package mysql

import (
	"A-Koya/react-PTJcist/adapter/repository"
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type mysqlHandler struct {
	db *sql.DB
}

func NewMysqlHandler() (*mysqlHandler, error) {
	db, err := sql.Open("mysql", "user1:password@tcp(localhost:3306)/cist_ptj")
	if err != nil {
		return &mysqlHandler{}, nil
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("DB ping failed")
		return &mysqlHandler{}, nil
	}
	defer db.Close()
	return &mysqlHandler{db: db}, nil
}

func (m *mysqlHandler) ExecuteContext(ctx context.Context, query string, args ...interface{}) error {
	_, err := m.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}
func (m *mysqlHandler) QueryContext(ctx context.Context, query string, args ...interface{}) (repository.Rows, error) {
	rows, err := m.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	row := newMysqlRows(rows)
	return row, nil
}
func (m *mysqlHandler) QueryRowContext(ctx context.Context, query string, args ...interface{}) repository.Row {
	row := m.db.QueryRowContext(ctx, query, args...)
	return newMysqlRow(row)
}
