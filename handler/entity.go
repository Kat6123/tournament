package handler

type (
	message struct {
		Content string `json:"message"`
	}

	balance struct {
		PlayerID int     `json:"playerId"`
		Balance  float32 `json:"balance"`
	}
)
