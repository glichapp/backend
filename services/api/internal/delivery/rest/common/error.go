package common

import "github.com/gin-gonic/gin"

type ErrorHandler func(ctx *gin.Context, err *APIError)

type APIError struct {
	StatusCode int
	ErrorPayload
}

type ErrorPayload struct {
	Error string `json:"error"`
}
