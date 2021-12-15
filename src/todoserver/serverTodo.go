package todoserver

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"todoApp/todoapi"
)

var apitodo *todoapi.ApiToDo

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
