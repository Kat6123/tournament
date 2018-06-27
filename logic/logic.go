package logic

import (
	"fmt"

	"github.com/kat6123/tournament/model"
)

type (
	PlayerProvider interface {
		LoadPlayer(playerID int) (*model.Player, error)
		SavePlayer(p *model.Player) error
	}

	TourProvider interface {
		LoadTournament(tourID int) (*model.Tournament, error)
		SaveTournament(t *model.Tournament) error
		CreateTournament(t *model.Tournament) error
	}

	Service struct {
		pp PlayerProvider
		tp TourProvider
	}
)

// Builder type should used to init package with repository.
type Builder struct {
	PP PlayerProvider
	TP TourProvider
}

// New method uses Builder to tune package logic.
func New(b Builder) *Service {
	return &Service{
		pp: b.PP,
		tp: b.TP,
	}
}

// Take loads player from repository, takes points and saves it.
func (s *Service) Take(playerID int, points float32) error {
	player, err := s.pp.LoadPlayer(playerID)
	if err != nil {
		return fmt.Errorf("load player with id %d from db: %v", playerID, err)
	}

	if err := player.Take(points); err != nil {
		return fmt.Errorf("take points from player %d: %v", playerID, err)
	}

	if err := s.pp.SavePlayer(player); err != nil {
		return fmt.Errorf("save player %d in db: %v", playerID, err)
	}

	return nil
}

// Fund loads player from repository, funds points and saves it.
func (s *Service) Fund(playerID int, points float32) error {
	player, err := s.pp.LoadPlayer(playerID)
	if err != nil {
		return fmt.Errorf("load player with id %d from db: %v", playerID, err)
	}

	player.Fund(points)

	if err := s.pp.SavePlayer(player); err != nil {
		return fmt.Errorf("save player %d in db: %v", playerID, err)
	}

	return nil
}

// AnnounceTournament creates tour and saves it in db. If already exists than return an error.
func (s *Service) AnnounceTournament(tourID int, deposit float32) error {
	tour := model.NewTour(tourID, deposit)

	if err := s.tp.CreateTournament(tour); err != nil {
		return fmt.Errorf("insert tour %d in db: %v", tourID, err)
	}

	return nil
}

// JoinTournament joins player to tour and saves tour in db.
func (s *Service) JoinTournament(tourID, playerID int) error {
	tour, err := s.tp.LoadTournament(tourID)
	if err != nil {
		return fmt.Errorf("load tournament with id %d from db: %v", tourID, err)
	}

	player, err := s.pp.LoadPlayer(playerID)
	if err != nil {
		return fmt.Errorf("load player with id %d from db: %v", playerID, err)
	}

	if err := tour.Join(player); err != nil {
		return err
	}

	if err := s.tp.SaveTournament(tour); err != nil {
		return fmt.Errorf("save tour %d in db: %v", tourID, err)
	}

	return nil
}

// Balance loads player and returns it balance.
func (s *Service) Balance(playerID int) (float32, error) {
	player, err := s.pp.LoadPlayer(playerID)
	if err != nil {
		return 0, fmt.Errorf("load player with id %d from db: %v", playerID, err)
	}

	return player.Balance, nil
}

// ResultTournament loads tour and returns its winner.
func (s *Service) ResultTournament(tourID int) (*model.Winner, error) {
	tour, err := s.tp.LoadTournament(tourID)
	if err != nil {
		return nil, fmt.Errorf("load tournament with id %d from db: %v", tourID, err)
	}

	return &tour.Winner, nil
}

// EndTournament loads tour, ends it and saves in db.
func (s *Service) EndTournament(tourID int) error {
	tour, err := s.tp.LoadTournament(tourID)
	if err != nil {
		return fmt.Errorf("load tournament with id %d from db: %v", tourID, err)
	}

	tour.End()

	if err := s.tp.SaveTournament(tour); err != nil {
		return fmt.Errorf("save tour %d in db: %v", tourID, err)
	}

	return nil
}
