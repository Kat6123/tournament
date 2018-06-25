package api

import (
	"github.com/kat6123/tournament/model"
)

// Repository interface needs to easy replace ddb engines.
type Repository interface {
	Connect() error
	Close()
	LoadPlayer(int) (*model.Player, error)
	SavePlayer(*model.Player) error
	LoadTournament(int) (*model.Tournament, error)
	SaveTournament(*model.Tournament) error
	CreateTournament(t *model.Tournament) error
}
