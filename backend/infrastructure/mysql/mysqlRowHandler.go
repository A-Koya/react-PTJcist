package mysql

import (
	"database/sql"
)

type mysqlRow struct {
	row *sql.Row
}

func newMysqlRow(row *sql.Row) mysqlRow {
	return mysqlRow{row: row}
}

func (mr mysqlRow) Scan(dest ...interface{}) error {
	if err := mr.row.Scan(dest...); err != nil {
		return err
	}
	return nil
}

type mysqlRows struct {
	rows *sql.Rows
}

func newMysqlRows(rows *sql.Rows) mysqlRows {
	return mysqlRows{rows: rows}
}

func (mrs mysqlRows) Scan(dist ...interface{}) error {
	if err := mrs.rows.Scan(dist...); err != nil {
		return err
	}
	return nil
}

func (mrs mysqlRows) Next() bool {
	return mrs.rows.Next()
}

func (mrs mysqlRows) Err() error {
	return mrs.rows.Err()
}

func (mrs mysqlRows) Close() error {
	return mrs.rows.Close()
}
