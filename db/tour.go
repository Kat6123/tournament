package db

import "github.com/kat6123/tournament/model"

func (tc *Tournaments) LoadTournament(tourID int) (*model.Tournament, error) {
	t := new(model.Tournament)
	err := tc.FindId(tourID).One(t)

	return t, err
}

func (tc *Tournaments) SaveTournament(t *model.Tournament) error {
	return tc.UpdateId(t.ID, t)
}

func (tc *Tournaments) CreateTournament(t *model.Tournament) error {
	return tc.Insert(t)
}

func (tc *Tournaments) DeleteTournament(tourID int) error {
	return tc.RemoveId(tourID)
}
