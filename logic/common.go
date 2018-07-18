package logic

import (
	"github.com/kat6123/tournament/errors"
)

func logicError(err error, msg string, entity string, id int) error {
	e := &errors.Error{
		ID:     id,
		Entity: entity,
		Err:    err,
	}

	t, ok := err.(*errors.Error)
	if ok {
		e.Kind = t.Kind
	} else {
		e.Kind = errors.Other
	}

	return errors.Wrap(e, msg)
}
