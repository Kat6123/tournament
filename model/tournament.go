package model

import (
	"fmt"
	"math/rand"
	"time"
)

// Tournament reflects tour ID, deposit, current balance and participants of tour.
// When tour has ended the winner field is set and balance is reset.
type Tournament struct {
	ID           int       `json:"id" bson:"_id"`
	Deposit      float32   `json:"deposit" bson:"deposit"`
	Balance      float32   `json:"balance,omitempty" bson:"balance,omitempty"`
	Participants []*Player `json:"players,omitempty" bson:"players,omitempty"`
	Winner       Winner    `json:"winners,omitempty" bson:"winners,omitempty"`
}

// Winner contains player and its prize.
type Winner struct {
	*Player `json:"player" bson:"player"`
	Prize   float32 `json:"prize" bson:"prize"`
}

// NewTour create new Tournament based on id and deposit.
func NewTour(id int, deposit float32) *Tournament {
	return &Tournament{
		ID:      id,
		Deposit: deposit,
	}
}

// End ends a tour if it hasn't been ended yet.
// It generates random player number and sets him as winner.
// Prize for winner is set to deposit and player balance increase on deposit. Balance is reset.
func (t *Tournament) End() {
	if !t.hasEnded() {
		n := t.generateWinnerNum()

		t.Winner = Winner{
			Player: t.Participants[n],
			Prize:  t.Deposit,
		}

		t.Participants[n].Fund(t.Deposit)
		t.Balance = 0
	}
}

// Join joins player to tour.
// If tour was ended, player is in tour already or  doesn't have enough funds returns an error.
func (t *Tournament) Join(p *Player) error {
	if t.hasEnded() {
		return fmt.Errorf("tournir %d was ended", t.ID)
	}

	if t.playerInTour(p.ID) {
		return fmt.Errorf("player %d already in tour %d", p.ID, t.ID)
	}

	if err := p.Take(t.Deposit); err != nil {
		return fmt.Errorf("take player %d funds when join to tour %d: %v", p.ID, t.ID, err)
	}

	t.Participants = append(t.Participants, p)
	t.Balance += t.Deposit

	return nil
}

func (t *Tournament) String() string {
	return fmt.Sprintf("Tournamnet:\nid %d \ndeposit %g \nbalance %g ", t.ID, t.Deposit, t.Balance)
}

func (t *Tournament) playerInTour(id int) bool {
	for i := range t.Participants {
		if id == t.Participants[i].ID {
			return true
		}
	}

	return false
}

func (t *Tournament) hasEnded() bool {
	return t.Winner != (Winner{})
}

func (t *Tournament) generateWinnerNum() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(len(t.Participants))
}
