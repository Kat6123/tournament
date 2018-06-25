package model

import (
	"fmt"
)

type Tournament struct {
	ID           int       `json:"id" bson:"_id"`
	Deposit      float32   `json:"deposit" bson:"deposit"`
	Participants []*Player `json:"players,omitempty" bson:"players,omitempty"`
	Balance      float32   `json:"balance,omitempty" bson:"balance,omitempty"`
	Winner       Winner    `json:"winners,omitempty" bson:"winners,omitempty"`
}

type Winner struct {
	*Player
	Prize float32
}

func (t *Tournament) SetDeposit(deposit float32) {
	t.Deposit = deposit
}

func (t *Tournament) SetWinner(num int) {
	// XXX: Check for current tour
	if &t.Winner != nil {
		t.Winner = Winner{
			Player: t.Participants[num],
			Prize:  t.Deposit,
		}

		t.Participants[num].Balance += t.Deposit

		t.Balance = 0
	}
}

func (t *Tournament) Join(p *Player) {
	// XXX
	if !t.playerInTour(p.ID) {
		t.Participants = append(t.Participants, p)
		p.Take(t.Deposit)
		t.Balance += t.Deposit
	}
}

func (t *Tournament) String() string {
	return fmt.Sprintf("Tournamnet:\n id %d \n deposit %g \n balance %g ", t.ID, t.Deposit, t.Balance)
}

func (t *Tournament) playerInTour(id int) bool {
	for i := range t.Participants {
		if id == t.Participants[i].ID {
			return true
		}
	}

	return false
}
