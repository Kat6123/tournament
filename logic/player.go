package logic

import "fmt"

// Take loads player from repository, takes points and saves it.
func Take(playerID int, points float32) error {
	player, err := ps.LoadPlayer(playerID)
	if err != nil {
		return fmt.Errorf("load player with id %d from db: %v", playerID, err)
	}

	if err := player.Take(points); err != nil {
		return fmt.Errorf("take points from player %d: %v", playerID, err)
	}

	if err := ps.SavePlayer(player); err != nil {
		return fmt.Errorf("save player %d in db: %v", playerID, err)
	}

	return nil
}

// Fund loads player from repository, funds points and saves it.
func Fund(playerID int, points float32) error {
	player, err := ps.LoadPlayer(playerID)
	if err != nil {
		return fmt.Errorf("load player with id %d from db: %v", playerID, err)
	}

	player.Fund(points)

	if err := ps.SavePlayer(player); err != nil {
		return fmt.Errorf("save player %d in db: %v", playerID, err)
	}

	return nil
}

// Balance loads player and returns it balance.
func Balance(playerID int) (float32, error) {
	player, err := ps.LoadPlayer(playerID)
	if err != nil {
		return 0, fmt.Errorf("load player with id %d from db: %v", playerID, err)
	}

	return player.Balance, nil
}
