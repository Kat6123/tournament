package db

import "github.com/kat6123/tournament/model"

// LoadTournament loads tour by id.
func (tc *Tournaments) LoadTournament(tourID int) (*model.Tournament, error) {
	t := new(model.Tournament)
	err := tc.FindId(tourID).One(t)

	return t, err
}

// SaveTournament update tour by id.
func (tc *Tournaments) SaveTournament(t *model.Tournament) error {
	return tc.UpdateId(t.ID, t)
}

// CreateTournament inserts new tour.
func (tc *Tournaments) CreateTournament(t *model.Tournament) error {
	return tc.Insert(t)
}

// DeleteTournament deletes by id.
func (tc *Tournaments) DeleteTournament(tourID int) error {
	return tc.RemoveId(tourID)
}
