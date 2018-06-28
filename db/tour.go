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

	return t, err
}

// Save update tour by id.
func (tc *Tournaments) Save(t *model.Tournament) error {
	return tc.UpdateId(t.ID, t)
}

// Create inserts new tour.
func (tc *Tournaments) Create(t *model.Tournament) error {
	return tc.Insert(t)
}

// delete deletes by id.
func (tc *Tournaments) delete(tourID int) error {
	return tc.RemoveId(tourID)
}
