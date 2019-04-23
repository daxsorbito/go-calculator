// Code generated by go-swagger; DO NOT EDIT.

package calc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/daxsorbito/go-calculator/models"
)

// CalcOperationOKCode is the HTTP code returned for type CalcOperationOK
const CalcOperationOKCode int = 200

/*CalcOperationOK OK

swagger:response calcOperationOK
*/
type CalcOperationOK struct {

	/*
	  In: Body
	*/
	Payload float32 `json:"body,omitempty"`
}

// NewCalcOperationOK creates CalcOperationOK with default headers values
func NewCalcOperationOK() *CalcOperationOK {

	return &CalcOperationOK{}
}

// WithPayload adds the payload to the calc operation o k response
func (o *CalcOperationOK) WithPayload(payload float32) *CalcOperationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the calc operation o k response
func (o *CalcOperationOK) SetPayload(payload float32) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CalcOperationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*CalcOperationDefault generic error

swagger:response calcOperationDefault
*/
type CalcOperationDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCalcOperationDefault creates CalcOperationDefault with default headers values
func NewCalcOperationDefault(code int) *CalcOperationDefault {
	if code <= 0 {
		code = 500
	}

	return &CalcOperationDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the calc operation default response
func (o *CalcOperationDefault) WithStatusCode(code int) *CalcOperationDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the calc operation default response
func (o *CalcOperationDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the calc operation default response
func (o *CalcOperationDefault) WithPayload(payload *models.Error) *CalcOperationDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the calc operation default response
func (o *CalcOperationDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CalcOperationDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
