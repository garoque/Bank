package err

import (
	"fmt"
)

// Error custom type err
type Error struct {
	HTTPCode int    `mapstructure:"code" json:"http_code,omitempty"`
	Message  string `mapstructure:"message" json:"message,omitempty"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %v - message: %v", e.HTTPCode, e.Message)
}

// New create a new err
func New(httpCode int, message string) error {
	return &Error{
		HTTPCode: httpCode,
		Message:  message,
	}
}
