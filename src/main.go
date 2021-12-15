package main

import (
	"net/http"
)

func getAllTodo(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/gettodo", getAllTodo)
	http.ListenAndServe(":8090", nil)
}
