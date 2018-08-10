// Package db provides types Players and Tournaments which allows work with MongoDB collections.
package db

import (
	"github.com/globalsign/mgo"
)

// New dials with db and init Players and Tournaments with collections,
// which are named 'players' and 'tours' accordingly.
func New(URL, tours, players string) (*Players, *Tournaments, error) {
	s, err := mgo.Dial(URL)
	if err != nil {
		return nil, nil, err
	}

	return &Players{s.DB("").C(players)},
		&Tournaments{s.DB("").C(tours)},
		nil
}
