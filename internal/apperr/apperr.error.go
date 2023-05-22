package apperr

import "fmt"

type CustomError struct {
	funcName,
	message string
	prevErr error
}

func NewCustomErr(funcName, message string, prevErr error) *CustomError {
	return &CustomError{
		funcName: funcName,
		message:  message,
		prevErr:  prevErr,
	}
}

func NewShortCustomErr(funcName string, prevErr error) *CustomError {
	return &CustomError{
		funcName: funcName,
		message:  "GRPC error",
		prevErr:  prevErr,
	}
}

func (err CustomError) Error() string {
	if err.prevErr != nil {
		return fmt.Sprintf("%s: %v. Function: %s", err.message, err.prevErr, err.funcName)
	}

	if err.prevErr != nil && err.message == "" {
		return fmt.Sprintf("%v. Function: %s", err.prevErr, err.funcName)
	}

	return fmt.Sprintf("%s: %v. Function: %s", err.message, err.prevErr, err.funcName)
}
