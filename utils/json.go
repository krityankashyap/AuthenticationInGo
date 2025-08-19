package utils

import (
	"encoding/json"
	"net/http"
  "github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func init(){
  Validator = NewValidator()

}

func NewValidator() *validator.Validate {
	// initialize the validator package
	return validator.New(validator.WithRequiredStructEnabled())

	
}


func WriteJsonResponse(w http.ResponseWriter, status int, data any) error {
   w.Header().Set("Content-Type" , "application/json")  // set the content type to application/json

	 w.WriteHeader(status)  // WriteHeader sends an HTTP response header with the provided status code.

	 return json.NewEncoder(w).Encode(data)  // NewEncoder returns a new encoder that writes to w. Encode writes the JSON encoding of v to the stream
}

func WriteJsonSuccessResponse(w http.ResponseWriter, status int, message string, data any) error {
   response := map[string]any {}

	 response["status"]= "success"
	 response["message"]= message
	 response["data"]= data
	 
	 return WriteJsonResponse(w, status, response)
}

func WriteJsonErrorResponse(w http.ResponseWriter, status int, message string, err error) error{ 
	response := map[string]any {}

	response["status"]= "error"
	response["message"]= message
	response["error"]= err

	return WriteJsonResponse(w, status, response)

}

func Readjson(r http.Request , result any) error {
	decoder := json.NewDecoder(r.Body)

	decoder.DisallowUnknownFields() // Prevents unknown fields from being included in the json body

	return decoder.Decode(result)
}