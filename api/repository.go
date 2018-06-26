package api

import (
	"github.com/kat6123/tournament/model"
)

// Repository interface needs to easy replace db engines.
type Repository interface {
	// Connect connects to repository.
	Connect() error
	// Close closes the connection.
	Close()
	// LoadPlayer loads player by ID.
	LoadPlayer(int) (*model.Player, error)
	// SavePlayer saved player.
	SavePlayer(*model.Player) error
	// LoadTournament loads tour by ID.
	LoadTournament(int) (*model.Tournament, error)
	// SaveTournament saves tour.
	SaveTournament(*model.Tournament) error
	// CreateTournament creates a new tour if it doesn't exist.
	CreateTournament(t *model.Tournament) error
}
