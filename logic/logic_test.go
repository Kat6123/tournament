package logic

import (
	"testing"

	"github.com/kat6123/tournament/logic/mocks"
	"github.com/kat6123/tournament/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_TakeNormal(t *testing.T) {
	playerID := 1
	points := float32(300)

	pp := new(mocks.PlayerProvider)
	pp.On("ByID", playerID).Return(&model.Player{ID: playerID, Balance: 300}, nil)
	pp.On("Save", mock.Anything).Return(nil)

	service := Service{pp: pp}
	err := service.Take(playerID, points)
	assert.Nil(t, err)

	pp.AssertExpectations(t)
}

func TestService_TakeErr(t *testing.T) {
	playerID := 1
	points := float32(300)

	pp := new(mocks.PlayerProvider)
	pp.On("ByID", playerID).Return(&model.Player{ID: playerID, Balance: 300}, nil)
	pp.On("Save", mock.Anything).Return(nil)

	service := Service{pp: pp}
	err := service.Take(playerID, points+100)
	assert.EqualError(t, err, "take points from player 1: insufficient funds")
	// Test that load was called?
}

//func TestService_Take(t *testing.T) {
//	const (
//		points = 300
//	)
//	pp := new(mocks.PlayerProvider)
//	pp.On("ByID", 1).Return(nil, fmt.Errorf("d"))
//	pp.On("ByID", 1).Return(nil, fmt.Errorf("d"))
//	pp.On("ByID", 1).Return(nil, fmt.Errorf("d"))
//	pp.On("Save", test.ID).Return(fmt.Errorf(test.savePlayerErr))
//
//	tt := []struct {
//		name string
//
//		ID            int
//		byIDErr       string
//		savePlayerErr string
//		expectedErr   string
//	}{
//		{
//			name:        "load error",
//			ID:          1,
//			byIDErr:     "load failed",
//			expectedErr: "load player with id 1 from db: load failed",
//		},
//		{
//			name:          "save error",
//			ID:           t 2,
//			savePlayerErr: "save failed",
//			expectedErr:   "take points player with id 1 from db: ?",
//		},
//	}
//
//	for _, tc := range t {
//		t.Run(test.name, func(t *testing.T) {
//			service := Service{pp: pp}
//			err := service.Take(test.ID, points)
//			assert.EqualError(t, err, test.expectedErr)
//
//			//pp.AssertExpectations(t)
//		})
//	}
//}
