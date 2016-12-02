package service

import (
    "fmt"
    "strings"
)

// Service error structure.
type ServiceError struct {
    error
    Message string
    Code    int
}

func (s *ServiceError) Error() string {
    return s.Message
}

func NewServiceError(err error, message string, code int, args ...string) *ServiceError {
    return &ServiceError{
        error: err,
        Code: code,
        Message: fmt.Sprintf("{\"error\": \"%s\", \"details\": [\"%s\"]}", message, strings.Join(args, "\", \"")),
    }
}



