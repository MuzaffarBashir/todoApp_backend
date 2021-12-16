package todoserver

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"todoApp/todoapi"
)

var apitodo *todoapi.ApiToDo

func (todoserver *ToDoServer) CreateToDo(request *http.Request, response http.ResponseWriter, conn *sql.DB) http.ResponseWriter {

	apitodo = todoapi.NewApiService() // getting instance of api layer
	apitodo, err := apitodo.CreateToDoApi(request, conn)
	if err != nil {
		response = getresponseSetting(response)
		response.WriteHeader(http.StatusBadRequest)
		return response
	} else {
		response = getresponseSetting(response)
		response.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(response).Encode(apitodo); err != nil {
			panic(err)
		}
	}
	return response

}

func (todoserver *ToDoServer) GetAllToDo(response http.ResponseWriter, connection *sql.DB) http.ResponseWriter {

	apitodo, err := apitodo.GetAllToDo(connection)
	if err != nil {

		response = getresponseSetting(response)
		response.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(response).Encode(nil); err != nil {
			panic(err)
		}
		return response
	} else {
		response = getresponseSetting(response)
		response.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(response).Encode(apitodo); err != nil {
			panic(err)
		}
	}
	return response

}
func getresponseSetting(response http.ResponseWriter) http.ResponseWriter {

	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	return response
}
