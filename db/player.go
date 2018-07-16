package db

import (
	"github.com/globalsign/mgo"
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
		return p, ConstructErr(err, "player", playerID)
	}

	return p, nil
}

// Save updates player model by id.
func (pc *Players) Save(p *model.Player) error {
	err := pc.UpdateId(p.ID, &p)
	if err != nil {
		return ConstructErr(err, "player", p.ID)
	}

	return nil
}

// delete deletes player by id.
func (pc *Players) delete(playerID int) error {
	err := pc.RemoveId(playerID)
	if err != nil {
		return ConstructErr(err, "player", playerID)
	}

	return nil
}

// create inserts new player.
func (pc *Players) create(p *model.Player) error {
	err := pc.Insert(p)
	if err != nil {
		return ConstructErr(err, "player", p.ID)
	}

	return nil
}
