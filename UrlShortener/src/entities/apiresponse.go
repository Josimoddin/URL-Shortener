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

// APIResponseInt Structure used to handle integer response
type ApiResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
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

// ThrowJSONResponse function is used to throw response in JSON format
func ThrowJSONResponse(response APIResponse, w http.ResponseWriter) {
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
func JSONParseErrorResponse() APIResponse {
	var response = APIResponse{}
	response.Status = false
	response.Message = "501 JSON parse Error."
	return response
}
