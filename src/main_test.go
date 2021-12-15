package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"todoApp/todoapi"

	"github.com/stretchr/testify/assert"
)

func TestGetAllTodo(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8090/gettodo", nil)
	response := httptest.NewRecorder()
	getAllTodo(response, request)
	res := response.Result()

	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	var todolist []todoapi.ApiToDo

	err := json.Unmarshal(data, &todolist)

	assert.Nil(t, err)
	assert.NotNil(t, todolist)
	assert.EqualValues(t, "new todo", todolist[0].Description)
	assert.EqualValues(t, http.StatusOK, res.StatusCode)

}
