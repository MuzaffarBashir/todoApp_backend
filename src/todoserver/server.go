package todoserver

import (
	"database/sql"
	"net/http"
)

type ServerTODO interface {
	GetAllToDo(response http.ResponseWriter, connection *sql.DB) http.ResponseWriter
	CreateToDo(r *http.Request, connection *sql.DB) http.ResponseWriter
}

type ToDoServer struct {
	Connection *sql.DB
}

func NewServerTODO() *ToDoServer {

	return &ToDoServer{}
}
