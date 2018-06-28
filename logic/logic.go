package logic

import (
	"fmt"

	"github.com/kat6123/tournament/model"
)

type (
	// PlayerProvider should be provided to use service of logic package. It should work with players collection.
	PlayerProvider interface {
		ByID(playerID int) (*model.Player, error)
		Save(p *model.Player) error
	}

	// TourProvider should be provided to use service of logic package. It should work with tours collection.
	TourProvider interface {
		ByID(tourID int) (*model.Tournament, error)
		Save(t *model.Tournament) error
		Create(t *model.Tournament) error
	}

	// Service provides API of logic package.
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

// Take ByIDs player from repository, takes points and saves it.
func (s *Service) Take(playerID int, points float32) error {
	player, err := s.pp.ByID(playerID)
	if err != nil {
		return fmt.Errorf("ByID player with id %d from db: %v", playerID, err)
	}

	if err := player.Take(points); err != nil {
		return fmt.Errorf("take points from player %d: %v", playerID, err)
	}

	if err := s.pp.Save(player); err != nil {
		return fmt.Errorf("save player %d in db: %v", playerID, err)
	}

	return nil
}

// Fund ByIDs player from repository, funds points and saves it.
func (s *Service) Fund(playerID int, points float32) error {
	player, err := s.pp.ByID(playerID)
	if err != nil {
		return fmt.Errorf("ByID player with id %d from db: %v", playerID, err)
	}

	player.Fund(points)

	if err := s.pp.Save(player); err != nil {
		return fmt.Errorf("save player %d in db: %v", playerID, err)
	}

	return nil
}

// AnnounceTournament creates tour and saves it in db. If already exists than return an error.
func (s *Service) AnnounceTournament(tourID int, deposit float32) error {
	tour := model.NewTour(tourID, deposit)

	if err := s.tp.Create(tour); err != nil {
		return fmt.Errorf("insert tour %d in db: %v", tourID, err)
	}

	return nil
}

// JoinTournament joins player to tour and saves tour in db.
func (s *Service) JoinTournament(tourID, playerID int) error {
	tour, err := s.tp.ByID(tourID)
	if err != nil {
		return fmt.Errorf("ByID tournament with id %d from db: %v", tourID, err)
	}

	player, err := s.pp.ByID(playerID)
	if err != nil {
		return fmt.Errorf("ByID player with id %d from db: %v", playerID, err)
	}

	if err := tour.Join(player); err != nil {
		return err
	}

	if err := s.tp.Save(tour); err != nil {
		return fmt.Errorf("save tour %d in db: %v", tourID, err)
	}

	return nil
}

// Balance ByIDs player and returns it balance.
func (s *Service) Balance(playerID int) (float32, error) {
	player, err := s.pp.ByID(playerID)
	if err != nil {
		return 0, fmt.Errorf("ByID player with id %d from db: %v", playerID, err)
	}

	return player.Balance, nil
}

// ResultTournament ByIDs tour and returns its winner.
func (s *Service) ResultTournament(tourID int) (*model.Winner, error) {
	tour, err := s.tp.ByID(tourID)
	if err != nil {
		return nil, fmt.Errorf("ByID tournament with id %d from db: %v", tourID, err)
	}

	return &tour.Winner, nil
}

// EndTournament ByIDs tour, ends it and saves in db.
func (s *Service) EndTournament(tourID int) error {
	tour, err := s.tp.ByID(tourID)
	if err != nil {
		return fmt.Errorf("ByID tournament with id %d from db: %v", tourID, err)
	}

	tour.End()

	if err := s.tp.Save(tour); err != nil {
		return fmt.Errorf("save tour %d in db: %v", tourID, err)
	}

	return nil
}
