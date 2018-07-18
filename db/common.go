package db

import (
	"fmt"

	"github.com/globalsign/mgo"
	"github.com/kat6123/tournament/errors"
)

func dbError(err error) error {
	var kind errors.Kind

	switch {
	case err == mgo.ErrNotFound:
		kind = errors.NotFound
	case mgo.IsDup(err):
		kind = errors.Duplicate
	default:
		kind = errors.Other
	}

	return &errors.Error{
		Kind: kind,
		Err:  fmt.Errorf("database: %v", err),
	}
}
