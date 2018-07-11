package errors

import (
	"fmt"
	"net/http"
)

const (
	// NotFound error.
	NotFound Kind = iota
	// Duplicate error.
	Duplicate
	// Other is the kind of not classified error.
	Other
)

type (
	// Kind of error.
	Kind int

	// TourError is a special type of error which is used by tour microservice.
	TourError struct {
		ID     int
		Kind   Kind
		Entity string
		Err    error
	}
)

func (e *TourError) Error() string {
	return fmt.Sprintf("%s %d: %s", e.Entity, e.ID, e.Err)
}

// Wrap return new error which was wrapped by s string.
// If e is of type TourError then Err field will be wrapped.
func Wrap(e error, s string) error {
	t, ok := e.(*TourError)
	if ok {
		t.Err = fmt.Errorf("%s: %v", s, t.Err)
		return t
	}

	return fmt.Errorf("%s: %v", s, e)

}

// HTTPStatus return status based on error type.
func HTTPStatus(e error) int {
	t, ok := e.(*TourError)
	if ok {
		switch t.Kind {
		case NotFound:
			return http.StatusNotFound
		case Duplicate:
			return http.StatusBadRequest
		default:
			return http.StatusInternalServerError
		}
	}

	return http.StatusInternalServerError
}
