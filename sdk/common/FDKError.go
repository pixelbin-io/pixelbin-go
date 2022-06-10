package common

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/stretchr/objx"
)

// FDKError is custom error type to pass more information along with the native error.
type FDKError struct {
	// Message is user-friendly message that can be displayed on front-end for end-users
	Message string `json:"message"`

	// Status is HTTP status code that is set in API response
	Status int `json:"status"`

	// A Code should be unique and used only once throughout the project
	// so that it can lead to the exact line of failed function when searched in the entire project.
	Code string `json:"code"`

	// Exception is technical expection title. This should be mostly 2-3 words title of the
	// error/er to be understood by developers.
	Exception string `json:"exception"`

	// Info is link/URL to the logging system(Sentry/Kibana) to get more insight and trace error
	Info string `json:"info"`

	// RequestID is the unique ID of a particular request for tracing the errors in the logging system(Sentry/Kibana)
	RequestID string `json:"request_id,omitempty"`

	//StackTrace gives the report of the active stack frames at a certain point in time during the execution of a program used for debugging purposes
	StackTrace string `json:"stack_trace,omitempty"`

	//Meta can be used for your custom keys, which would be helpful for detailed error and for debugging
	Meta objx.Map `json:"meta,omitempty"`
}

// NewFDKError constructs and returns new FDKError object
func NewFDKError(message string) *FDKError {
	if message == "" {
		message = errors.New("Something went wrong [sdk]").Error()
	}
	return &FDKError{
		Message:    message,
		Status:     http.StatusInternalServerError,
		Code:       "",
		Exception:  "",
		Info:       "",
		RequestID:  "",
		StackTrace: "",
		Meta:       objx.Map{},
	}
}

// Error returns technical error message with error code and er title
func (f *FDKError) Error() string {
	return fmt.Sprintf("%s", f.Message)
}

func (f *FDKError) String() string {
	return f.Error()
}

// SetStatus sets the HTTP Status in the error object
func (f *FDKError) SetStatus(httpStatus int) *FDKError {
	if httpStatus < 200 || httpStatus > 600 {
		return f
	}
	f.Status = httpStatus
	return f
}

// SetRequestID sets the RequestID in the error object
func (f *FDKError) SetRequestID(requestID string) *FDKError {
	f.RequestID = requestID
	return f
}
