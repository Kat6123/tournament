package db

import (
	"github.com/globalsign/mgo"
	"github.com/kat6123/tournament/errors"
)

func constructErr(e error, et string, id int) error {
	var kind errors.Kind

	switch {
	case e == mgo.ErrNotFound:
		kind = errors.NotFound
	case mgo.IsDup(e):
		kind = errors.Duplicate
	default:
		kind = errors.Other
	}

	return &errors.TourError{
		ID:     id,
		Entity: et,
		Kind:   kind,
		Err:    e,
	}
}
