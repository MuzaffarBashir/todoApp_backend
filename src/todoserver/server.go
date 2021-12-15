package todoserver

import (
	"database/sql"
	"net/http"
)

type ServerTODO interface {
	GetTODO(response http.ResponseWriter, connection *sql.DB) http.ResponseWriter
}

type ToDoServer struct {
	Connection *sql.DB
}

func NewServerTODO() *ToDoServer {

	return &ToDoServer{}
}
