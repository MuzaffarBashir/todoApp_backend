package tests

import (
	"database/sql"
	"log"
	"todoApp/tododb"

	"github.com/DATA-DOG/go-sqlmock"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	tododb.SetConnection(db)
	if err != nil {
		log.Fatalf("wasn't thinking of err '%s' while geting connection", err)
	}
	return db, mock
}
