package todoserver

import (
	"net/http"
	"testing"
	"todoApp/tododb"

	"github.com/stretchr/testify/assert"
)

// Success case of getting all todos
func TestGetAllToDo(t *testing.T) {
	db := tododb.GetConnection()
	servertodo := &ToDoServer{

		Connection: db,
	}
	var response http.ResponseWriter
	response = servertodo.GetAllToDo(response, servertodo.Connection)
	//assert.Nil(t, err)
	assert.NotNil(t, response)

}
