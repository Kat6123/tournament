package db

import (
	"github.com/globalsign/mgo"
	"github.com/kat6123/tournament/model"
)

// Tournaments provides basic methods to work with collection of model.Tournament.
type Tournaments struct {
	*mgo.Collection
}

// ByID loads tour by id.
func (tc *Tournaments) ByID(tourID int) (*model.Tournament, error) {
	t := new(model.Tournament)
	err := tc.FindId(tourID).One(t)
	if err != nil {
		return t, dbError(err)
	}

	return t, nil
}

// Save update tour by id.
func (tc *Tournaments) Save(t *model.Tournament) error {
	err := tc.UpdateId(t.ID, t)
	if err != nil {
		return dbError(err)
	}

	return nil
}

// Create inserts new tour.
func (tc *Tournaments) Create(t *model.Tournament) error {
	err := tc.Insert(t)
	if err != nil {
		return dbError(err)
	}

	return nil
}

// delete deletes by id.
func (tc *Tournaments) delete(tourID int) error {
	err := tc.RemoveId(tourID)
	if err != nil {
		return dbError(err)
	}

	return nil
}
