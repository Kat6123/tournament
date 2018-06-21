package logic

import (
	"math/rand"
	"time"

	"github.com/kat6123/tournament/db"
	"github.com/kat6123/tournament/model"
)

func Take(playerId int, points float32) error {
	player, err := db.LoadPlayer(playerId)
	if err != nil {
		return err
	}

	player.Take(points)

	if err := db.SavePlayer(player); err != nil {
		return err
	}
	return nil
}

func Fund(playerId int, points float32) error {
	player, err := db.LoadPlayer(playerId)
	if err != nil {
		return err
	}

	player.Fund(points)

	if err := db.SavePlayer(player); err != nil {
		return err
	}
	return nil
}

func AnnounceTournament(tourId int, deposit float32) error {
	tour := &model.Tournament{
		ID:      tourId,
		Deposit: deposit,
	}

	err := db.CreateTournament(tour)
	return err
}

func JoinTournament(tourId, playerId int) error {
	tour, err := db.LoadTournament(tourId)
	if err != nil {
		return err
	}

	player, err := db.LoadPlayer(playerId)
	if err != nil {
		return err
	}

	tour.Join(player)

	if err := db.SaveTournament(tour); err != nil {
		return err
	}

	return nil
}

func Balance(playerId int) (*model.Player, error) {
	return db.LoadPlayer(playerId)
}

func ResultTournament(tourId int) (*model.Winner, error) {
	tour, err := db.LoadTournament(tourId)
	if err != nil {
		return nil, err
	}

	return &tour.Winner, nil
}

func EndTournament(tourId int) error {
	tour, err := db.LoadTournament(tourId)
	if err != nil {
		return err
	}

	rand.Seed(time.Now().UnixNano())
	winN := rand.Intn(len(tour.Participants))
	tour.SetWinner(winN)

	if err := db.SaveTournament(tour); err != nil {
		return err
	}

	return nil
}
