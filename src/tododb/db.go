package tododb

import "database/sql"

type DbServices interface {
	GetAllToDO(conn *sql.DB) (*Todo, error)
}
type Todo struct {
	Description string
	ID          string
	Connection  *sql.DB
}

func NewTodo() *Todo {
	return &Todo{}
}
