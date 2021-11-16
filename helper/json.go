package helper

import (
	"base-project/app/model/response"
	"encoding/json"
	"net/http"
)

func JsonWriter(writer http.ResponseWriter,code int, status string, data interface{}) {
	resp := response.Standard{
		Code:   code,
		Status: status,
		Data:   data,
	}
	err := json.NewEncoder(writer).Encode(resp)
	Panic(err)
}
