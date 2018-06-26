package model

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTournament_hasEnded(t *testing.T) {
	tt := []struct {
		name           string
		tour           Tournament
		expectedResult bool
	}{
		{
			name: "tour is going",
			tour: Tournament{
				ID:      123,
				Deposit: 234,
			},
			expectedResult: false,
		},
		{
			name: "tour was ended and winner is set",
			tour: Tournament{
				ID:      123,
				Deposit: 234,
				Winner: Winner{
					Player: &Player{ID: 123, Balance: 123},
					Prize:  123,
				},
			},
			expectedResult: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.tour.hasEnded(), tc.expectedResult)
		})
	}
}

// How to test random int?
func TestTournament_generateWinnerNum(t *testing.T) {
	tcNum := 5
	sameCount := 0
	prev := 0

	tour := Tournament{
		ID:           123,
		Deposit:      123,
		Participants: make([]*Player, 5),
	}

	for i := 0; i < tcNum; i++ {
		got := tour.generateWinnerNum()
		fmt.Println(got)
		if got == prev {
			sameCount++
		}

		if got < 0 || got > len(tour.Participants) {
			t.Fatalf("generated winner is out of range: got %d, limit %d", got, len(tour.Participants))
		}
		prev = got
	}

	if sameCount == tcNum-1 {
		t.Errorf("get the same winners during all tests: test number %d", len(tour.Participants))
	}
}

func TestTournament_playerInTour(t *testing.T) {
	tour := Tournament{
		ID:      123,
		Deposit: 123,
		Participants: []*Player{
			{
				ID:      1,
				Balance: 235,
			},
			{
				ID:      2,
				Balance: 235,
			},
		},
	}

	tt := []struct {
		name        string
		ID          int
		expectedRes bool
	}{
		{
			name:        "player in tour",
			ID:          1,
			expectedRes: true,
		},
		{
			name:        "player not in tour",
			ID:          123,
			expectedRes: false,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.expectedRes, tour.playerInTour(tc.ID), "test that %s", tc.name)
	}
}

func TestTournament_String(t *testing.T) {
	tc := struct {
		name        string
		tour        Tournament
		expectedStr string
	}{
		name:        "increase balance",
		tour:        Tournament{ID: 123, Deposit: 300, Balance: 600},
		expectedStr: "Tournamnet:\nid 123 \ndeposit 300 \nbalance 600 ",
	}

	assert.Equal(t, tc.expectedStr, tc.tour.String())
}

func TestTournament_Join(t *testing.T) {
	tt := []struct {
		name                                       string
		tour                                       Tournament
		player                                     *Player
		expectedPlayerInTour                       bool
		expectedTourBalance, expectedPlayerBalance float32
		expectedErrStr                             string
	}{
		{
			name:                  "simple joining",
			tour:                  Tournament{ID: 123, Deposit: 300},
			player:                &Player{ID: 1, Balance: 400},
			expectedPlayerInTour:  true,
			expectedTourBalance:   300,
			expectedPlayerBalance: 100,
		},
		{
			name:                  "player already in tour",
			tour:                  Tournament{ID: 123, Deposit: 300, Participants: []*Player{{ID: 1, Balance: 400}}},
			player:                &Player{ID: 1, Balance: 400},
			expectedPlayerInTour:  true,
			expectedTourBalance:   0,
			expectedPlayerBalance: 400,
			expectedErrStr:        "player 1 already in tour 123",
		},
		{
			name: "tour was ended",
			tour: Tournament{
				ID:      123,
				Deposit: 300,
				Winner: Winner{
					Player: &Player{ID: 123},
				}},
			player:                &Player{ID: 1, Balance: 400},
			expectedPlayerInTour:  false,
			expectedTourBalance:   0,
			expectedPlayerBalance: 400,
			expectedErrStr:        "tournir 123 was ended",
		},
		{
			name:                  "player doesn't have enough money",
			tour:                  Tournament{ID: 123, Deposit: 300},
			player:                &Player{ID: 1, Balance: 200},
			expectedPlayerInTour:  false,
			expectedTourBalance:   0,
			expectedPlayerBalance: 200,
			expectedErrStr:        "take player 1 funds when join to tour 123: insufficient funds",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.tour.Join(tc.player)

			if err != nil {
				require.EqualError(t, err, tc.expectedErrStr, "errors doesn't match")
			}

			assert.Equal(t, tc.player.Balance, tc.expectedPlayerBalance, "player balance doesn't match")
			assert.Equal(t, tc.tour.Balance, tc.expectedTourBalance, "tour balance doesn't match")
			assert.Equal(t, tc.expectedPlayerInTour, tc.tour.playerInTour(tc.player.ID),
				"player in tour return wrong result")
		})
	}
}

// Can be better
func TestTournament_End(t *testing.T) {
	tt := []struct {
		name                                       string
		tour                                       Tournament
		expectedWinnerPrize, expectedWinnerBalance float32
		expectedWinnerID                           int
	}{
		{
			name: "end tour",
			tour: Tournament{
				ID: 123, Balance: 300, Deposit: 300, Participants: []*Player{{ID: 1, Balance: 400}}},
			expectedWinnerID:      1,
			expectedWinnerPrize:   300,
			expectedWinnerBalance: 700,
		},
		{
			name: "tour was ended",
			tour: Tournament{
				ID: 123, Deposit: 300,
				Participants: []*Player{{ID: 1, Balance: 400}},
				Winner: Winner{
					Player: &Player{ID: 1, Balance: 700},
					Prize:  300,
				}},
			expectedWinnerID:      1,
			expectedWinnerPrize:   300,
			expectedWinnerBalance: 700,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tc.tour.End()

			assert.Equal(t, tc.expectedWinnerBalance, tc.tour.Winner.Balance, "winner balance doesn't match")
			assert.Equal(t, tc.expectedWinnerID, tc.tour.Winner.ID, "winner id doesn't match")
			assert.Equal(t, tc.expectedWinnerPrize, tc.tour.Winner.Prize, "winner prize doesn't match")
			assert.Equal(t, float32(0), tc.tour.Balance, "tour balance not equal to 0")
		})
	}
}

func TestNewTour(t *testing.T) {
	tour := Tournament{
		ID:      123,
		Deposit: 123,
	}

	assert.Equal(t, tour, *NewTour(123, 123), "new tour id and deposit doesn't match")
}
