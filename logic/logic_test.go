package logic

import (
	"testing"

	"fmt"

	"github.com/kat6123/tournament/logic/mocks"
	"github.com/kat6123/tournament/model"
	"github.com/stretchr/testify/assert"
)

func TestService_Take(t *testing.T) {
	const points = 200
	tt := []struct {
		name           string
		ID             int
		expectedErrMsg string
	}{
		{
			name: "normal",
			ID:   1,
		},
		{
			name:           "load error",
			ID:             2,
			expectedErrMsg: "load player with id 2 from db: load failed",
		},
		{
			name:           "take error",
			ID:             3,
			expectedErrMsg: "take points from player 3: insufficient funds",
		},
		{
			name:           "save error",
			ID:             4,
			expectedErrMsg: "save player 4 in db: save failed",
		},
	}

	pp := new(mocks.PlayerProvider)
	pp.On("ByID", 1).Return(&model.Player{ID: 1, Balance: 300}, nil)
	pp.On("Save", &model.Player{ID: 1, Balance: 300 - points}).Return(nil)

	pp.On("ByID", 2).Return(nil, fmt.Errorf("load failed"))

	// Balance less than points to cause take error
	pp.On("ByID", 3).Return(&model.Player{ID: 3, Balance: 100}, nil)

	pp.On("ByID", 4).Return(&model.Player{ID: 4, Balance: 300}, nil)
	pp.On("Save", &model.Player{ID: 4, Balance: 300 - points}).Return(
		fmt.Errorf("save failed"))

	service := Service{pp: pp}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := service.Take(tc.ID, points)
			if err != nil {
				assert.EqualError(t, err, tc.expectedErrMsg)
			}
		})
	}

	pp.AssertExpectations(t)
}

func TestService_Fund(t *testing.T) {
	const points = 200
	tt := []struct {
		name           string
		ID             int
		expectedErrMsg string
	}{
		{
			name: "normal",
			ID:   1,
		},
		{
			name:           "load error",
			ID:             2,
			expectedErrMsg: "load player with id 2 from db: load failed",
		},
		{
			name:           "save error",
			ID:             3,
			expectedErrMsg: "save player 3 in db: save failed",
		},
	}

	pp := new(mocks.PlayerProvider)
	pp.On("ByID", 1).Return(&model.Player{ID: 1, Balance: 300}, nil)
	pp.On("Save", &model.Player{ID: 1, Balance: 300 + points}).Return(nil)

	pp.On("ByID", 2).Return(nil, fmt.Errorf("load failed"))

	pp.On("ByID", 3).Return(&model.Player{ID: 3, Balance: 300}, nil)
	pp.On("Save", &model.Player{ID: 3, Balance: 300 + points}).Return(
		fmt.Errorf("save failed"))

	service := Service{pp: pp}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := service.Fund(tc.ID, points)
			if err != nil {
				assert.EqualError(t, err, tc.expectedErrMsg)
			}
		})
	}

	pp.AssertExpectations(t)
}

func TestService_AnnounceTournament(t *testing.T) {
	tt := []struct {
		name             string
		ID               int
		deposit          float32
		expectedErrorMsg string
	}{
		{
			name:    "normal",
			ID:      1,
			deposit: 300,
		},
		{
			name:             "create error",
			ID:               2,
			deposit:          300,
			expectedErrorMsg: "insert tour 2 in db: create failed",
		},
	}

	tp := new(mocks.TourProvider)
	tp.On("Create", &model.Tournament{ID: tt[0].ID, Deposit: tt[0].deposit}).Return(nil)
	tp.On("Create", &model.Tournament{ID: tt[1].ID, Deposit: tt[1].deposit}).Return(
		fmt.Errorf("create failed"))

	service := Service{tp: tp}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := service.AnnounceTournament(tc.ID, tc.deposit)
			if err != nil {
				assert.EqualError(t, err, tc.expectedErrorMsg)
			}
		})
	}
}

func TestService_JoinTournament(t *testing.T) {
	tt := []struct {
		name             string
		tourID           int
		playerID         int
		expectedErrorMsg string
	}{
		{
			name:     "normal",
			tourID:   1,
			playerID: 1,
		},
		{
			name:             "load tour error",
			tourID:           2,
			playerID:         2,
			expectedErrorMsg: "load tournament with id 2 from db: load failed",
		},
		{
			name:             "load player error",
			tourID:           3,
			playerID:         3,
			expectedErrorMsg: "load player with id 3 from db: load failed",
		},
		{
			name:             "join error",
			tourID:           4,
			playerID:         4,
			expectedErrorMsg: "tournir 4 was ended",
		},
		{
			name:             "save error",
			tourID:           5,
			playerID:         5,
			expectedErrorMsg: "save tour 5 in db: save failed",
		},
	}

	pp := new(mocks.PlayerProvider)
	pp.On("ByID", tt[0].playerID).Return(&model.Player{ID: tt[0].playerID}, nil)

	pp.On("ByID", tt[2].playerID).Return(
		nil, fmt.Errorf("load failed"))

	pp.On("ByID", tt[3].playerID).Return(&model.Player{ID: tt[3].playerID}, nil)

	pp.On("ByID", tt[4].playerID).Return(&model.Player{ID: tt[4].playerID}, nil)

	tp := new(mocks.TourProvider)
	tp.On("ByID", tt[0].tourID).Return(&model.Tournament{ID: tt[0].tourID}, nil)
	tp.On("Save",
		&model.Tournament{ID: tt[0].tourID, Participants: []*model.Player{{ID: tt[0].playerID}}}).Return(nil)

	tp.On("ByID", tt[1].tourID).Return(
		&model.Tournament{}, fmt.Errorf("load failed"))

	tp.On("ByID", tt[2].tourID).Return(nil, nil)

	// Return tour that was ended.
	tp.On("ByID", tt[3].tourID).Return(
		&model.Tournament{ID: tt[3].tourID, Winner: model.Winner{Prize: 100}}, nil)

	tp.On("ByID", tt[4].tourID).Return(&model.Tournament{ID: tt[4].tourID}, nil)
	tp.On("Save",
		&model.Tournament{ID: tt[4].tourID, Participants: []*model.Player{{ID: tt[4].playerID}}}).Return(
		fmt.Errorf("save failed"))

	s := Service{pp: pp, tp: tp}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := s.JoinTournament(tc.tourID, tc.playerID)
			if err != nil {
				assert.EqualError(t, err, tc.expectedErrorMsg)
			}
		})
	}
}

func TestService_Balance(t *testing.T) {
	pp := new(mocks.PlayerProvider)
	pp.On("ByID", 1).Return(&model.Player{ID: 1, Balance: 300}, nil)
	pp.On("ByID", 2).Return(nil, fmt.Errorf("load failed"))

	tt := []struct {
		name           string
		ID             int
		expectedErrMsg string
	}{
		{
			name: "normal",
			ID:   1,
		},
		{
			name:           "load error",
			ID:             2,
			expectedErrMsg: "load player 2 from db: load failed",
		},
	}

	service := Service{pp: pp}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, err := service.Balance(tc.ID)
			if err != nil {
				assert.EqualError(t, err, tc.expectedErrMsg)
			}
		})
	}

	pp.AssertExpectations(t)
}

func TestService_ResultTournament(t *testing.T) {
	tp := new(mocks.TourProvider)
	tp.On("ByID", 1).Return(&model.Tournament{}, nil)
	tp.On("ByID", 2).Return(nil, fmt.Errorf("load failed"))

	tt := []struct {
		name           string
		ID             int
		expectedErrMsg string
	}{
		{
			name: "normal",
			ID:   1,
		},
		{
			name:           "load error",
			ID:             2,
			expectedErrMsg: "laod tournament 2 from db: load failed",
		},
	}

	service := Service{tp: tp}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, err := service.ResultTournament(tc.ID)
			if err != nil {
				assert.EqualError(t, err, tc.expectedErrMsg)
			}
		})
	}

	tp.AssertExpectations(t)
}

//func TestService_EndTournament(t *testing.T) {
//	tp := new(mocks.TourProvider)
//	tp.On("ByID", 1).Return(
//		&model.Tournament{ID: 1, Participants: []*model.Player{{ID: 1}}}, nil)
//	tp.On("Save",
//		&model.Tournament{ID: 1, Participants: []*model.Player{{ID: 1}}, Winner: model.Winner{Player: &model.Player{ID: 1}}}).Return(
//		&model.Tournament{ID: 1}, nil)
//	tp.On("ByID", 2).Return(nil, fmt.Errorf("load failed"))
//
//	tt := []struct {
//		name           string
//		ID             int
//		expectedErrMsg string
//	}{
//		{
//			name: "normal",
//			ID:   1,
//		},
//		{
//			name:           "load error",
//			ID:             2,
//			expectedErrMsg: "load tournament 2 from db: load failed",
//		},
//		{
//			name:           "save error",
//			ID:             3,
//			expectedErrMsg: "save tournament 3 from db: load failed",
//		},
//	}
//
//	service := Service{tp: tp}
//	for _, tc := range tt {
//		t.Run(tc.name, func(t *testing.T) {
//			err := service.EndTournament(tc.ID)
//			if err != nil {
//				assert.EqualError(t, err, tc.expectedErrMsg)
//			}
//		})
//	}
//
//	tp.AssertExpectations(t)
//}
