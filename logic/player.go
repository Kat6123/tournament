package logic

import "fmt"

type Player struct {
	ID      int     `json:"id"`
	Balance float32 `json:"balance"`
}

func GetPlayer(PlayerId int) *Player {
	// load from db
	return &Player{
		ID:      PlayerId,
		Balance: 300,
	}
}

func (p *Player) Take(points float32) {
	p.Balance -= points
	// save in db
}

func (p *Player) Fund(points float32) {
	p.Balance += points
	// save in db
}

func (p *Player) String() string {
	return fmt.Sprintf("Player:\n id %d \n balance %g", p.ID, p.Balance)
}
