package models

import "fmt"

type ErrInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *ErrInfo) Error() string {
	return fmt.Sprintf("Error code: %s, Error message: %s", e.Code, e.Message)
}
