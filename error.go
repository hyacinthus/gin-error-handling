package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Error predefined error message
var (
	ErrInvalidID = NewError(400, "InvalidID", "Invalid id in request url")
	ErrNoAuth    = NewError(401, "NoAuth", "Authentication required]")
	ErrForbidden = NewError(403, "Forbidden", "It's not your resource")
)

// Error custom http error type
type Error struct {
	// http status code
	Code int `json:"-"`
	// a key for client sorting errors
	Key string `json:"error"`
	// more readable error message
	Msg string `json:"msg"`
}

// NewError new a custom error type pointer
func NewError(code int, key string, msg string) *Error {
	return &Error{
		Code: code,
		Key:  key,
		Msg:  msg,
	}
}

// Error makes it compatible with `error` interface.
func (e *Error) Error() string {
	return e.Key + ": " + e.Msg
}

// Handle handle errors in gin handler, intelligently generate response body with right http status code.
func Handle(c *gin.Context, err error) {
	var parsedError *Error
	switch err.(type) {
	case *Error:
		parsedError = err.(*Error)
	default:
		if err == ErrNotFound {
			// catch 404 errors, you should try every data store provider's "not found error".
			parsedError = NewError(http.StatusNotFound, "NotFound", err.Error())
		} else {
			// return real server error message in debug mode only
			debug := true // fake debug config
			if debug {
				parsedError = NewError(http.StatusInternalServerError, "ServerError", err.Error())
			} else {
				parsedError = NewError(http.StatusInternalServerError, "ServerError", "Internal Server Error")
			}
		}
	}
	// write json error body and abort gin context chain
	c.AbortWithStatusJSON(parsedError.Code, parsedError)
}
