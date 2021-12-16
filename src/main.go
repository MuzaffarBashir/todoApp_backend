package main

import (
	"database/sql"
	"net/http"
	"todoApp/tododb"
	"todoApp/todoserver"
)

var Connection *sql.DB
var servertodo *todoserver.ToDoServer

func init() {
	Connection = tododb.GetConnection()
	servertodo = todoserver.NewServerTODO()
}

func getAllTodo(w http.ResponseWriter, r *http.Request) {
	// calling to get list of todos
	servertodo.GetAllToDo(w, Connection)
}
func handlerequest(response http.ResponseWriter, request *http.Request) {
	//calling server for requestion validation to create TODO

	servertodo.CreateToDo(request, response, Connection)
}

func main() {
	http.HandleFunc("/gettodo", getAllTodo)
	http.HandleFunc("/", handlerequest)
	http.ListenAndServe(":8090", nil)
}
