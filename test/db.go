package test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

// GetDB return a new mock instance of DB
func GetDB() (*sqlx.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual), sqlmock.MonitorPingsOption(true))
	return sqlx.NewDb(db, "mysql"), mock
}

// NewRows return model to add rows
func NewRows(columns ...string) *sqlmock.Rows {
	return sqlmock.NewRows(columns)
}
