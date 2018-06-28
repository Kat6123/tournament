package db

import "github.com/kat6123/tournament/model"

// LoadPlayer loads player by id.
func (pc *Players) LoadPlayer(playerID int) (*model.Player, error) {
	p := new(model.Player)
	err := pc.FindId(playerID).One(p)

	return p, err
}

// Player will be inited even if err?

// SavePlayer updates player model by id.
func (pc *Players) SavePlayer(p *model.Player) error {
	return pc.UpdateId(p.ID, &p)
}

// DeletePlayer deletes player by id.
func (pc *Players) DeletePlayer(playerID int) error {
	return pc.RemoveId(playerID)
}

// CreatePlayer inserts new player.
func (pc *Players) CreatePlayer(p *model.Player) error {
	return pc.Insert(p)
}

// Drop drops 'players' collection.
func (pc *Players) Drop() error {
	return pc.DropCollection()
}
