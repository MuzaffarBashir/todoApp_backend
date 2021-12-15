package main

import (
	"database/sql"
	"net/http"
	"todoApp/todoserver"
)

var Connection *sql.DB
var servertodo *todoserver.ToDoServer

func getAllTodo(w http.ResponseWriter, r *http.Request) {
	// calling to get list of todos
	servertodo.GetAllToDo(w, Connection)
}

func main() {
	http.HandleFunc("/gettodo", getAllTodo)
	http.ListenAndServe(":8090", nil)
}
