package model

import "fmt"

type Player struct {
	ID      int     `json:"playerId" bson:"_id"`
	Balance float32 `json:"balance" bson:"balance"`
}

func (p *Player) Take(points float32) {
	p.Balance -= points
}

func (p *Player) Fund(points float32) {
	p.Balance += points
}

func (p *Player) String() string {
	return fmt.Sprintf("Player:\n id %d \n balance %g", p.ID, p.Balance)
}
