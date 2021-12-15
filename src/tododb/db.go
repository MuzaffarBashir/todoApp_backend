package tododb

import (
	"database/sql"
	"fmt"
)

const (
	DBhost     = "localhost"
	DBport     = 5432
	DBuser     = "postgres"
	DBpassword = "Abc123"
	Dbname     = "userdata"
)

var Connection *sql.DB
var err error

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
func GetConnection() *sql.DB {
	psqlcon := fmt.Sprintf("host= %s port= %d user= %s password= %s dbname= %s sslmode=disable",
		DBhost, DBport, DBuser, DBpassword, Dbname)

	Connection, err = sql.Open("postgres", psqlcon)

	if err != nil {

		fmt.Print("Error in getting Coonection", err)
		panic(err.Error())

	} else {
		fmt.Println("connected")
	}
	return Connection

}
