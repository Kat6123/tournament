package db

import "github.com/kat6123/tournament/model"

// Player will be inited even if err
func (pc *Players) LoadPlayer(playerID int) (*model.Player, error) {
	p := new(model.Player)
	err := pc.FindId(playerID).One(p)

	return p, err
}

func (pc *Players) SavePlayer(p *model.Player) error {
	return pc.UpdateId(p.ID, &p)
}

func (pc *Players) DeletePlayer(playerID int) error {
	return pc.RemoveId(playerID)
}

func (pc *Players) CreatePlayer(p *model.Player) error {
	return pc.Insert(p)
}
