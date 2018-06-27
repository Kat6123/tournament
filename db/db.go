package db

import (
	"github.com/globalsign/mgo"
)

type (
	// Players implements PlayerProvider interface.
	Players struct {
		*mgo.Collection
	}

	// Tournaments implements TournamentProvider interface and.
	Tournaments struct {
		*mgo.Collection
	}
)

// New dials with db and init Players and Tournaments with Mongo collections,
// which are named 'players' and 'tours' accordingly.
func New(URL string, DB string) (*Players, *Tournaments, error) {
	s, err := mgo.Dial(URL)
	if err != nil {
		return nil, nil, err
	}

	return &Players{s.DB(DB).C("players")},
		&Tournaments{s.DB(DB).C("tours")},
		nil
}

// How to close the session?
