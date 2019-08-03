package utils

import (
	"encoding/json"
	"net/http"
)

type CreateResponse struct {
	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`

	StatusCode int
	
	ErrorMessage string
}

// NewCreateResponse creates CreateResponse with default header values 
func NewCreateResponse() *CreateResponse {
	return &CreateResponse{}
}
func NewCreateResponseWithParams(status int, payload interface{}, errMsg string) *CreateResponse {
	createResponse := new(CreateResponse)
	if status != 0 {
		createResponse.StatusCode = status
	}
	if payload != nil {
		createResponse.Payload = payload
	}
	if len(errMsg) != 0 {
		createResponse.ErrorMessage = errMsg
	}
	return createResponse
}

// WithPayload adds the payload to the create task response
func (o *CreateResponse) WithPayload(payload interface{}) *CreateResponse {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create task response
func (o *CreateResponse) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WithStatusCode adds the status to the create task response
func (o *CreateResponse) WithStatusCode(code int) *CreateResponse {
	o.StatusCode = code
	return o
}

// SetStatusCode sets the status to the create task response
func (o *CreateResponse) SetStatusCode(code int) {
	o.StatusCode = code
}

// WithErrorMessage sets the error message to the create task response
func (o *CreateResponse) WithErrorMessage(errMessage string) *CreateResponse {
	o.ErrorMessage = errMessage
	return o
}

// SetErrorMessage sets the error message to the create task response
func (o *CreateResponse) SetErrorMessage(errMessage string) {
	o.ErrorMessage = errMessage
}

// WriteResponse to the client
func (o *CreateResponse) WriteResponse(w http.ResponseWriter) {
	if o.StatusCode == 0 {
		if o.Payload != nil {
			o.StatusCode = http.StatusOK
		} else {
			o.StatusCode = http.StatusInternalServerError
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(o.StatusCode)
	if o.Payload != nil {
		payload := o.Payload
		json.NewEncoder(w).Encode(payload)
	} else {
		json.NewEncoder(w).Encode(o.ErrorMessage)
	}
}