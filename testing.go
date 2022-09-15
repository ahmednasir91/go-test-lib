package testing

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

// MockDb is a wrapper around sqlmock and sqlx.DB
type MockDb struct {
	mockDB *sql.DB
	Mock   sqlmock.Sqlmock
	Db     *sqlx.DB
}

// Finish is a wrapper around sqlmock.ExpectationsWereMet
func (d *MockDb) Finish() {
	err := d.mockDB.Close()
	if err != nil {
		return
	}
}

// NewSqlxMockDB creates a new MockDb
func NewSqlxMockDB() *MockDb {
	mockDB, mock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	return &MockDb{mockDB, mock, sqlxDB}
}