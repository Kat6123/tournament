package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayer_Fund(t *testing.T) {
	tc := struct {
		name            string
		balance, points float32
		expectedBalance float32
	}{
		name:            "increase balance",
		balance:         300,
		points:          100,
		expectedBalance: 400,
	}

	p := Player{ID: 123, Balance: tc.balance}
	p.Fund(tc.points)

	assert.Equal(t, tc.expectedBalance, p.Balance)
}

func TestPlayer_Take(t *testing.T) {
	tt := []struct {
		name              string
		balance, points   float32
		expectedBalance   float32
		expectedErrString string
	}{
		{
			name:            "decrease balance",
			balance:         300,
			points:          100,
			expectedBalance: 200,
		},
		{
			name:              "decrease to negative balance",
			balance:           20,
			points:            100,
			expectedBalance:   20,
			expectedErrString: "insufficient funds",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := Player{ID: 123, Balance: tc.balance}
			err := p.Take(tc.points)

			assert.Equal(t, tc.expectedBalance, p.Balance)
			if err != nil {
				assert.EqualError(t, err, tc.expectedErrString)
			}
		})
	}
}

func TestPlayer_String(t *testing.T) {
	tc := struct {
		name        string
		player      Player
		expectedStr string
	}{
		name:        "increase balance",
		player:      Player{ID: 123, Balance: 300},
		expectedStr: "Player:\nid 123 \nbalance 300",
	}

	assert.Equal(t, tc.expectedStr, tc.player.String())
}
