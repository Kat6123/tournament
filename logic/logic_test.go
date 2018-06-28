package logic

import (
	"testing"

	"github.com/kat6123/tournament/mocks"
	"github.com/kat6123/tournament/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_Take(t *testing.T) {
	tt := []struct {
		name            string
		balance         float32
		points          float32
		expectedBalance float32
		expectedErrMsg  string
	}{
		{
			name:            "take 100 points from 200",
			balance:         200,
			points:          100,
			expectedBalance: 100,
		},
		{
			name:            "insufficient funds",
			balance:         50,
			points:          100,
			expectedBalance: 50,
			expectedErrMsg:  "take points from player 1: insufficient funds",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			player := &model.Player{ID: 1, Balance: tc.balance}
			ppMock := new(mocks.PlayerProvider)
			ppMock.On("LoadPlayer", 1).Return(player, nil)
			ppMock.On("SavePlayer", mock.Anything).Return(nil)

			s := New(Builder{
				PP: ppMock,
			})

			err := s.Take(player.ID, tc.points)
			if err != nil {
				assert.EqualError(t, err, tc.expectedErrMsg, "error should be generated")
				ppMock.AssertNotCalled(t, "SavePlayer")
			} else {
				ppMock.AssertExpectations(t)
			}
			assert.Equal(t, player.Balance, tc.expectedBalance, "balance doesn't match")

			//ppMock.AssertExpectations(t) Обязательно проверять?
		})
	}
}

func TestService_Fund(t *testing.T) {
	tt := []struct {
		name            string
		balance         float32
		points          float32
		expectedBalance float32
		expectedErrMsg  string
	}{
		{
			name:            "fund 100 points",
			balance:         200,
			points:          100,
			expectedBalance: 300,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			player := &model.Player{ID: 1, Balance: tc.balance}
			ppMock := new(mocks.PlayerProvider)
			ppMock.On("LoadPlayer", 1).Return(player, nil)
			ppMock.On("SavePlayer", mock.Anything).Return(nil)

			s := New(Builder{
				PP: ppMock,
			})

			err := s.Fund(player.ID, tc.points)
			if err != nil {
				assert.EqualError(t, err, tc.expectedErrMsg, "error should be generated")
				ppMock.AssertNotCalled(t, "SavePlayer")
			} else {
				ppMock.AssertExpectations(t)
			}
			assert.Equal(t, player.Balance, tc.expectedBalance, "balance doesn't match")

			//ppMock.AssertExpectations(t) Обязательно проверять?
		})
	}
}

// Не нужно проверять что и как инициализировалось внутри?
//func TestService_AnnounceTournament(t *testing.T) {
//	tt := []struct {
//		name           string
//		tourID         int
//		deposit        float32
//		expectedErrMsg string
//	}{
//		{
//			name:   "announce tour with ID and deposit",
//			tourID: 123,
//			// How to check for duplicate and etc.?
//			// Check for negative deposit?
//			deposit: 200,
//		},
//	}
//
//	for _, tc := range tt {
//		t.Run(tc.name, func(t *testing.T) {
//			createdTour := new(model.Tournament)
//
//			tpMock := new(mocks.TourProvider)
//			tpMock.On("CreateTournament", createdTour).Return(
//				nil)
//
//			s := New(Builder{
//				TP: tpMock,
//			})
//
//			expectedTour := model.Tournament{
//				Deposit: tc.deposit,
//			}
//
//			_ = s.AnnounceTournament(tc.tourID, tc.deposit)
//			assert.Equal(t, createdTour, expectedTour)
//
//			tpMock.AssertExpectations(t)
//		})
//	}
//}
