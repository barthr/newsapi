package newsapi

import "fmt"

// Error defines an API error from newsapi.
type Error struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e Error) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// APIError returns if the given err is of type `newsapi.Error`.
func APIError(err error) bool {
	_, ok := err.(*Error)
	return ok
}
