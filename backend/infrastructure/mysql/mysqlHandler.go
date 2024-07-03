package mysql

import (
	"A-Koya/react-PTJcist/adapter/repository"
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type mysqlHandler struct {
	db *sql.DB
}

func NewMysqlHandler() (*mysqlHandler, error) {
	dsn := "user1:password@tcp(mysql:3306)/cist_ptj?charset=utf8mb4&parseTime=True&loc=Asia%2FTokyo"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(fmt.Printf("db connection failed:%e", err))
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(fmt.Printf("db ping failed:%e", err))
	}
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
