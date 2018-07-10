package db

import (
	"github.com/globalsign/mgo"
	"github.com/kat6123/tournament/errors"
	"github.com/kat6123/tournament/model"
)

// Players provides basic methods to work with collection of model.Player.
type Players struct {
	*mgo.Collection
}

// ByID loads player by id.
func (pc *Players) ByID(playerID int) (*model.Player, error) {
	p := new(model.Player)
	err := pc.FindId(playerID).One(p)

	if err != nil {
		errKind := errors.Other
		if err == mgo.ErrNotFound {
			errKind = errors.NotFound
		}

		return nil, &errors.Error{
			ID:     playerID,
			Entity: "player",
			Kind:   errKind,
			Err:    err,
		}
	}
	return p, err
}

// Save updates player model by id.
func (pc *Players) Save(p *model.Player) error {
	return pc.UpdateId(p.ID, &p)
}

// delete deletes player by id.
func (pc *Players) delete(playerID int) error {
	return pc.RemoveId(playerID)
}

// create inserts new player.
func (pc *Players) create(p *model.Player) error {
	return pc.Insert(p)
}
