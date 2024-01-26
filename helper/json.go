package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestbody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	errorEncoder := encoder.Encode(response)
	PanicIfError(errorEncoder)
}
