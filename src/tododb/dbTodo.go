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
func (todo *Todo) CreateToDo(Connection *sql.DB) (int64, error) {

	lastInsertedID := 0
	insert := `INSERT INTO "todolist" ("description") VALUES ($1) RETURNING ID`
	Connection.QueryRow(insert, todo.Description).Scan(&lastInsertedID)
	if lastInsertedID == 0 {
		err := errors.New("todo not inserted")
		return 0, err
	}
	return int64(lastInsertedID), nil
}
