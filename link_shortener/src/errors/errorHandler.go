package syserrors

import "fmt"

type ErrorHandler interface {
	InternalDAError(err error)
	AnalyticsServiceError(err error)
}

type defaultErrorHandler struct {
}

func CreateDefaultErrorHandler() ErrorHandler {
	return new(defaultErrorHandler)
}

func (handler *defaultErrorHandler) InternalDAError(err error) {
	fmt.Println(err.Error())
}

func (handler *defaultErrorHandler) AnalyticsServiceError(err error) {
	fmt.Printf("Error contacting analytics service %s\n", err.Error())
}
