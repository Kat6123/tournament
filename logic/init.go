package logic

import "github.com/kat6123/tournament/api"

var (
	ps api.PlayerService
	ts api.TournamentService
)

// Builder type should be used to init package with repository.
type Builder struct {
	PlayerService api.PlayerService
	TourService   api.TournamentService
}

// Init method uses Builder to tune package logic.
func Init(b Builder) {
	ps = b.PlayerService
	ts = b.TourService
}
