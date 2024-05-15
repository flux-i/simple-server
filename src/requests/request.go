package requests

import (
	"encoding/json"
	"log"
	"net/http"
)

// RequestError represents the parameter error in the HTTP request
type requestError struct {
	Location    string `json:"location"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Follow the same semantics of cornice
type errorResponse struct {
	Status string         `json:"status"`
	Errors []requestError `json:"errors"`
}

type successResponse struct {
	DataType string      `json:"type"`
	Result   interface{} `json:"result"`
}

// Request object represents the base application request
type Request struct {
	errors      []requestError
	request     *http.Request
	response    http.ResponseWriter
	errorStatus int
	payloadType string
	payloadData interface{}
}

// HasErrors returns true if there are any errors in the request
func (req *Request) HasErrors() bool {
	return len(req.errors) > 0
}

// AddError adds the errors that the request has encountered
func (req *Request) AddError(status int, location, name, desc string) {
	req.errorStatus = status
	err := requestError{Location: location, Name: name, Description: desc}
	req.errors = append(req.errors, err)
}

// WriteErrorResponse writes error response to the supplied
// http.ResponseWriter
func (req *Request) WriteErrorResponse() {
	req.response.Header().Set("Content-Type", "application/json")
	if req.errorStatus == 0 {
		req.response.WriteHeader(http.StatusInternalServerError)
	} else {
		req.response.WriteHeader(req.errorStatus)
	}
	if len(req.errors) <= 0 {
		return
	}
	resp := errorResponse{Status: "error", Errors: req.errors}
	response, err := json.Marshal(resp)
	if err != nil {
		// We can't do much about it here
		log.Println("Error while constructing json error response")
	} else {
		req.response.Write(response)
	}
}

// Valid will deserialize the data from the http request and set
// appropriate errors if there any errors on the data and return
// whether the validation of data was successful or not.
func (req *Request) Valid() bool {
	log.Println("Not Implemented: Concrete classes should implement Valid()")
	return false
}

// AddResponsePayload adds the payload data for success response
func (req *Request) AddResponsePayload(retType string, data interface{}) {
	req.payloadType = retType
	req.payloadData = data
}

// WriteSuccessResponse writes the success response and attaches
// payload if there are any.
func (req *Request) WriteSuccessResponse() {
	req.response.Header().Set("Content-Type", "application/json")
	req.response.WriteHeader(http.StatusOK)
	if req.payloadData == nil {
		return
	}
	resp := successResponse{DataType: req.payloadType, Result: req.payloadData}
	response, err := json.Marshal(resp)
	if err != nil {
		// We can't do much about it here
		log.Println("Error while construction json success response")
	} else {
		req.response.Write(response)
	}
}
