package logic

import (
	"fmt"

	"github.com/kat6123/tournament/model"
)

// AnnounceTournament creates tour and saves it in db. If already exists than return an error.
func AnnounceTournament(tourID int, deposit float32) error {
	tour := model.NewTour(tourID, deposit)

	if err := ts.CreateTournament(tour); err != nil {
		return fmt.Errorf("insert tour %d in db: %v", tourID, err)
	}

	return nil
}

// JoinTournament joins player to tour and saves tour in db.
func JoinTournament(tourID, playerID int) error {
	tour, err := ts.LoadTournament(tourID)
	if err != nil {
		return fmt.Errorf("load tournament with id %d from db: %v", tourID, err)
	}

	player, err := ps.LoadPlayer(playerID)
	if err != nil {
		return fmt.Errorf("load player with id %d from db: %v", playerID, err)
	}

	if err := tour.Join(player); err != nil {
		return err
	}

	if err := ts.SaveTournament(tour); err != nil {
		return fmt.Errorf("save tour %d in db: %v", tourID, err)
	}

	return nil
}

// ResultTournament loads tour and returns its winner.
func ResultTournament(tourID int) (*model.Winner, error) {
	tour, err := ts.LoadTournament(tourID)
	if err != nil {
		return nil, fmt.Errorf("load tournament with id %d from db: %v", tourID, err)
	}

	return &tour.Winner, nil
}

// EndTournament loads tour, ends it and saves in db.
func EndTournament(tourID int) error {
	tour, err := ts.LoadTournament(tourID)
	if err != nil {
		return fmt.Errorf("load tournament with id %d from db: %v", tourID, err)
	}

	tour.End()

	if err := ts.SaveTournament(tour); err != nil {
		return fmt.Errorf("save tour %d in db: %v", tourID, err)
	}

	return nil
}
