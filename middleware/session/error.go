package session

import "fmt"

type sessionError struct {
	prob string
}

func (e *sessionError) Error() string {
	return fmt.Sprintf("%s", e.prob)
}
