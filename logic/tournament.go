package logic

import "fmt"

type Tournament struct {
	ID           int     `json:"id" bson:"_id"`
	Deposit      float32 `json:"deposit" bson:"deposit"`
	participants []*Player
	balance      float32
	Winners      []*struct {
		Player
		prize float64
	}
}

func GetTournament(TournamentId int) *Tournament {
	// load from db
	return &Tournament{
		ID:      TournamentId,
		Deposit: 100,
	}
}

func (t *Tournament) SetDeposit(deposit float32) {
	t.Deposit = deposit
	// save in db
}

func (t *Tournament) Join(p *Player) {
	t.participants = append(t.participants, p)
	p.Take(t.Deposit)
	t.balance += t.Deposit
	// save in db
}

func (t *Tournament) String() string {
	return fmt.Sprintf("Tournamnet:\n id %d \n deposit %g \n balance %g ", t.ID, t.Deposit, t.balance)
}
