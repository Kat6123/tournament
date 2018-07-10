package errors

import (
	"fmt"
	"net/http"
)

const (
	NotFound kind = iota
	Duplicate
	Other
)

type (
	kind   int
	entity string

	Error struct {
		ID     int
		Kind   kind
		Entity entity
		Err    error
	}
)

func (e *Error) Error() string {
	return fmt.Sprintf("%s %d: %s", e.Entity, e.ID, e.Err)
}

func (e *Error) wrap(s string) {
	e.Err = fmt.Errorf("%s: %v", s, e.Err)
}

func (e *Error) httpStatus() int {
	switch e.Kind {
	case NotFound:
		return http.StatusNotFound
	case Duplicate:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
