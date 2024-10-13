package entities

import (
	"encoding/json"
	"log"
	"net/http"
)

// APIResponse Structure used to handle http response using json
type APIResponse struct {
	Status   bool   `json:"success"`
	Message  string `json:"message"`
	Response string `json:"response"`
}

// ErrorResponse Structure used to handle error  response using json
type ErrorResponse struct {
	Status   bool     `json:"success"`
	Message  string   `json:"message"`
	Response []string `json:"response"`
}

// BlankPathCheckResponse function is used to return blank path response
func BlankPathCheckResponse() APIResponse {
	var response = APIResponse{}
	response.Status = false
	response.Message = "404 not found."
	log.Println("Blank request called")
	return response
}

// NotPostMethodResponse function is used to return not post method response
func NotPostMethodResponse() APIResponse {
	var response = APIResponse{}
	response.Status = false
	response.Message = "405 method not allowed."
	return response
}

// InternalServerErrorResponse function is used to return Internal server error response
func InternalServerErrorResponse() APIResponse {
	var response = APIResponse{}
	response.Status = false
	response.Message = "Internal Server Error."
	return response
}

// JSONParseErrorResponse function is used to return JSON parse error response
func JSONParseErrorResponse() APIResponse {
	var response = APIResponse{}
	response.Status = false
	response.Message = "501 JSON parse Error."
	return response
}

// DbConErrorResponse function is used to return database connection error response
func DbConErrorResponse() APIResponse {
	var response = APIResponse{}
	response.Status = false
	response.Message = "502 DB Connection Error."
	return response
}

// DbErrorResponse function is used to return database Insertion error response
func DbErrorResponse(message string) APIResponse {
	var response = APIResponse{}
	response.Status = false
	response.Message = message
	return response
}

// GetParamsErrorResponse function is used to return parameters error in GET request error response
func GetParamsErrorReponse() APIResponse {
	var response = APIResponse{}
	response.Status = false
	response.Message = "422 Get Params Error."
	return response
}

// ThrowJSONResponse function is used to throw response in JSON format
func ThrowJSONResponse(response APIResponse, w http.ResponseWriter) {
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// ThrowErrorResponse function is used to throw response in JSON format
func ThrowErrorResponse(responseErr []string, w http.ResponseWriter) {
	var response = ErrorResponse{}
	response.Status = false
	response.Message = "201 Operational Error."
	response.Response = responseErr
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
