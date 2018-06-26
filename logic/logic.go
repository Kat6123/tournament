package logic

import (
	"fmt"

	"github.com/kat6123/tournament/api"
	"github.com/kat6123/tournament/model"
)

var db api.Repository

// Builder type should used to init package with repository.
type Builder struct {
	R api.Repository
}

// Init method uses Builder to tune package logic.
func Init(b Builder) {
	db = b.R
}

// Take loads player from repository, takes points and saves it.
func Take(playerID int, points float32) error {
	player, err := db.LoadPlayer(playerID)
	if err != nil {
		return fmt.Errorf("load player with id %d from db: %v", playerID, err)
	}

	if err := player.Take(points); err != nil {
		return fmt.Errorf("take points from player %d: %v", playerID, err)
	}

	if err := db.SavePlayer(player); err != nil {
		return fmt.Errorf("save player %d in db: %v", playerID, err)
	}

	return nil
}

// Fund loads player from repository, funds points and saves it.
func Fund(playerID int, points float32) error {
	player, err := db.LoadPlayer(playerID)
	if err != nil {
		return fmt.Errorf("load player with id %d from db: %v", playerID, err)
	}

	player.Fund(points)

	if err := db.SavePlayer(player); err != nil {
		return fmt.Errorf("save player %d in db: %v", playerID, err)
	}

	return nil
}

// AnnounceTournament creates tour and saves it in db. If already exists than return an error.
func AnnounceTournament(tourID int, deposit float32) error {
	tour := model.NewTour(tourID, deposit)

	if err := db.CreateTournament(tour); err != nil {
		return fmt.Errorf("insert tour %d in db: %v", tourID, err)
	}

	return nil
}

// JoinTournament joins player to tour and saves tour in db.
func JoinTournament(tourID, playerID int) error {
	tour, err := db.LoadTournament(tourID)
	if err != nil {
		return fmt.Errorf("load tournament with id %d from db: %v", tourID, err)
	}

	player, err := db.LoadPlayer(playerID)
	if err != nil {
		return fmt.Errorf("load player with id %d from db: %v", playerID, err)
	}

	if err := tour.Join(player); err != nil {
		return err
	}

	if err := db.SaveTournament(tour); err != nil {
		return fmt.Errorf("save tour %d in db: %v", tourID, err)
	}

	return nil
}

// Balance loads player and returns it balance.
func Balance(playerID int) (float32, error) {
	player, err := db.LoadPlayer(playerID)
	if err != nil {
		return 0, fmt.Errorf("load player with id %d from db: %v", playerID, err)
	}

	return player.Balance, nil
}

// ResultTournament loads tour and returns its winner.
func ResultTournament(tourID int) (*model.Winner, error) {
	tour, err := db.LoadTournament(tourID)
	if err != nil {
		return nil, fmt.Errorf("load tournament with id %d from db: %v", tourID, err)
	}

	return &tour.Winner, nil
}

// EndTournament loads tour, ends it and saves in db.
func EndTournament(tourID int) error {
	tour, err := db.LoadTournament(tourID)
	if err != nil {
		return fmt.Errorf("load tournament with id %d from db: %v", tourID, err)
	}

	tour.End()

	if err := db.SaveTournament(tour); err != nil {
		return fmt.Errorf("save tour %d in db: %v", tourID, err)
	}

	return nil
}
