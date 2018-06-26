package model

import "fmt"

// Player reflects its player ID and its balance.
type Player struct {
	ID      int     `json:"playerId" bson:"_id"`
	Balance float32 `json:"balance" bson:"balance"`
}

// Take decrease balance by points.
// If balance will be negative after points are taken then return error.
// Balance doesn't change in that case.
func (p *Player) Take(points float32) (err error) {
	if p.Balance-points < 0 {
		return fmt.Errorf("insufficient funds")
	}
	p.Balance -= points

	return
}

// Fund increase balance with points.
func (p *Player) Fund(points float32) {
	p.Balance += points
}

// String format player when print like:
// Player:
// id 123
// balance 563
func (p *Player) String() string {
	return fmt.Sprintf("Player:\nid %d \nbalance %g", p.ID, p.Balance)
}
