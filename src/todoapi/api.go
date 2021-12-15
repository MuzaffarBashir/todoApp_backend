package todoapi

import "database/sql"

type ApiService interface {
	Gettodo(connection *sql.DB) (*ApiToDo, error)
}
type ApiToDo struct {
	Description string `jason:"Description"`
	ID          string `jason:"ID"`
	Connection  *sql.DB
}

func NewApiService() *ApiToDo {
	return &ApiToDo{}
}
