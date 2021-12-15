package todoapi

import "database/sql"

type ApiService interface {
	GetAllToDo(connection *sql.DB) ([]ApiToDo, error)
}
type ApiToDo struct {
	Description string `jason:"Description"`
	ID          string `jason:"ID"`
	Connection  *sql.DB
}

func NewApiService() *ApiToDo {
	return &ApiToDo{}
}
