package logic

import (
	"math/rand"
	"time"

	"fmt"

	"github.com/kat6123/tournament/api"
	"github.com/kat6123/tournament/model"
)

var db api.Repository

type Builder struct {
	R api.Repository
}

func Init(b Builder) {
	db = b.R
}

func Take(playerID int, points float32) error {
	player, err := db.LoadPlayer(playerID)
	if err != nil {
		return fmt.Errorf("load player with id %d from db: %v", playerID, err)
	}

	player.Take(points)

	if err := db.SavePlayer(player); err != nil {
		return fmt.Errorf("save player %d in db: %v", playerID, err)
	}
	return nil
}

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

func AnnounceTournament(tourID int, deposit float32) error {
	tour := &model.Tournament{
		ID:      tourID,
		Deposit: deposit,
	}

	if err := db.CreateTournament(tour); err != nil {
		return fmt.Errorf("insert tour %d in db: %v", tourID, err)
	}
	return nil
}

func JoinTournament(tourID, playerID int) error {
	tour, err := db.LoadTournament(tourID)
	if err != nil {
		return fmt.Errorf("load tournament with id %d from db: %v", tourID, err)
	}

	player, err := db.LoadPlayer(playerID)
	if err != nil {
		return fmt.Errorf("load player with id %d from db: %v", playerID, err)
	}

	tour.Join(player)

	if err := db.SaveTournament(tour); err != nil {
		return fmt.Errorf("save tour %d in db: %v", tourID, err)
	}

	return nil
}

func Balance(playerID int) (float32, error) {
	player, err := db.LoadPlayer(playerID)
	if err != nil {
		return 0, fmt.Errorf("load player with id %d from db: %v", playerID, err)
	}

	return player.Balance, nil
}

func ResultTournament(tourID int) (*model.Winner, error) {
	tour, err := db.LoadTournament(tourID)
	if err != nil {
		return nil, fmt.Errorf("load tournament with id %d from db: %v", tourID, err)
	}

	return &tour.Winner, nil
}

func EndTournament(tourID int) error {
	tour, err := db.LoadTournament(tourID)
	if err != nil {
		return fmt.Errorf("load tournament with id %d from db: %v", tourID, err)
	}

	rand.Seed(time.Now().UnixNano())
	winN := rand.Intn(len(tour.Participants))
	tour.SetWinner(winN)

	if err := db.SaveTournament(tour); err != nil {
		return fmt.Errorf("save tour %d in db: %v", tourID, err)
	}

	return nil
}
