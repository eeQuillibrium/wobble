package errors

import (
	"go.uber.org/zap/zapcore"
	"net/http"
)

type httpStatus int

type CustomError struct {
	Message  string
	HttpCode httpStatus
	LogLevel zapcore.Level
}

func NewInternal() CustomError {
	return CustomError{
		Message:  "Internal server error",
		HttpCode: http.StatusInternalServerError,
		LogLevel: zapcore.ErrorLevel,
	}
}
