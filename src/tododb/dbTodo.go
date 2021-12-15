package tododb

import (
	"database/sql"
	"errors"
)

func (todo *Todo) GetAllToDO(conn *sql.DB) (*sql.Rows, error) {

	retrieve := `SELECT id,description FROM "todolist"`
	data, err := conn.Query(retrieve)
	if err != nil {
		err = errors.New("db error getting todos list")
		return nil, err
	}
	return data, nil
}
