package logic

import (
	"testing"

	"fmt"

	"github.com/kat6123/tournament/logic/mocks"
	"github.com/kat6123/tournament/model"
	"github.com/stretchr/testify/assert"
)

const points = 200

func TestService_Take(t *testing.T) {
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
			expectedErrMsg: "player 2: load: failed",
		},
		{
			name:           "take error",
			ID:             3,
			expectedErrMsg: "player 3: take points: insufficient funds",
		},
		{
			name:           "save error",
			ID:             4,
			expectedErrMsg: "player 4: save: failed",
		},
	}

	pp := new(mocks.PlayerProvider)
	defer pp.AssertExpectations(t)

	pp.On("ByID", 1).Return(&model.Player{ID: 1, Balance: 300}, nil)
	pp.On("Save", &model.Player{ID: 1, Balance: 100}).Return(nil)

	pp.On("ByID", 2).Return(nil,
		fmt.Errorf("failed"))

	// Balance less than points to cause take error
	pp.On("ByID", 3).Return(&model.Player{ID: 3, Balance: 100}, nil)

	pp.On("ByID", 4).Return(&model.Player{ID: 4, Balance: 300}, nil)
	pp.On("Save", &model.Player{ID: 4, Balance: 100}).Return(
		fmt.Errorf("failed"))

	service := Service{pp: pp}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := service.Take(tc.ID, points)
			if err != nil {
				assert.EqualError(t, err, tc.expectedErrMsg)
			}
		})
	}
}

func TestService_Fund(t *testing.T) {
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
			expectedErrMsg: "player 2: load: failed",
		},
		{
			name:           "save error",
			ID:             3,
			expectedErrMsg: "player 3: save: failed",
		},
	}

	pp := new(mocks.PlayerProvider)
	defer pp.AssertExpectations(t)

	pp.On("ByID", 1).Return(&model.Player{ID: 1, Balance: 300}, nil)
	pp.On("Save", &model.Player{ID: 1, Balance: 500}).Return(nil)

	pp.On("ByID", 2).Return(nil, fmt.Errorf("failed"))

	pp.On("ByID", 3).Return(&model.Player{ID: 3, Balance: 300}, nil)
	pp.On("Save", &model.Player{ID: 3, Balance: 500}).Return(
		fmt.Errorf("failed"))

	service := Service{pp: pp}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := service.Fund(tc.ID, points)
			if err != nil {
				assert.EqualError(t, err, tc.expectedErrMsg)
			}
		})
	}
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
			expectedErrorMsg: "tour 2: insert: failed",
		},
	}

	tp := new(mocks.TourProvider)
	defer tp.AssertExpectations(t)
	tp.On("Create", &model.Tournament{ID: 1, Deposit: 300}).Return(nil)
	tp.On("Create", &model.Tournament{ID: 2, Deposit: 300}).Return(
		fmt.Errorf("failed"))

	service := Service{tp: tp}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := service.AnnounceTournament(tc.ID, tc.deposit)
			if err != nil {
				fmt.Println(err)
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
			expectedErrorMsg: "tour 2: load: failed",
		},
		{
			name:             "load player error",
			tourID:           3,
			playerID:         3,
			expectedErrorMsg: "player 3: load: failed",
		},
		{
			name:             "join error",
			tourID:           4,
			playerID:         4,
			expectedErrorMsg: "player 4: join: tournir 4 was ended",
		},
		{
			name:             "save error",
			tourID:           5,
			playerID:         5,
			expectedErrorMsg: "player 5: load: failed",
		},
	}

	pp := new(mocks.PlayerProvider)
	defer pp.AssertExpectations(t)

	pp.On("ByID", 1).Return(&model.Player{ID: 1}, nil)

	pp.On("ByID", 3).Return(
		nil, fmt.Errorf("failed"))

	pp.On("ByID", 4).Return(&model.Player{ID: 4}, nil)

	pp.On("ByID", 5).Return(&model.Player{ID: 5}, nil)

	tp := new(mocks.TourProvider)
	defer tp.AssertExpectations(t)

	tp.On("ByID", 1).Return(&model.Tournament{ID: 1}, nil)
	tp.On("Save",
		&model.Tournament{ID: 1, Participants: []*model.Player{{ID: 1}}}).Return(nil)

	tp.On("ByID", 2).Return(
		&model.Tournament{}, fmt.Errorf("failed"))

	tp.On("ByID", 3).Return(nil, nil)

	// Return tour that was ended.
	tp.On("ByID", 4).Return(
		&model.Tournament{ID: 4, Winner: model.Winner{Prize: 100}}, nil)

	tp.On("ByID", 5).Return(&model.Tournament{ID: 5}, nil)
	tp.On("Save",
		&model.Tournament{ID: 5, Participants: []*model.Player{{ID: 5}}}).Return(
		fmt.Errorf("failed"))

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
	defer pp.AssertExpectations(t)

	pp.On("ByID", 1).Return(&model.Player{ID: 1, Balance: 300}, nil)
	pp.On("ByID", 2).Return(nil, fmt.Errorf("failed"))

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
			expectedErrMsg: "player 2: load: failed",
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
	defer tp.AssertExpectations(t)

	tp.On("ByID", 1).Return(&model.Tournament{}, nil)
	tp.On("ByID", 2).Return(nil, fmt.Errorf("failed"))

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
			expectedErrMsg: "tour 2: load: failed",
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
}

func TestService_EndTournament(t *testing.T) {
	tp := new(mocks.TourProvider)
	defer tp.AssertExpectations(t)

	tp.On("ByID", 1).Return(
		&model.Tournament{ID: 1, Participants: []*model.Player{{ID: 1}}}, nil)
	tp.On("Save",
		&model.Tournament{ID: 1, Participants: []*model.Player{{ID: 1}},
			Winner: model.Winner{Player: &model.Player{ID: 1}}}).Return(
		nil)

	tp.On("ByID", 2).Return(nil, fmt.Errorf("failed"))

	tp.On("ByID", 3).Return(
		&model.Tournament{ID: 3, Participants: []*model.Player{{ID: 3}}}, nil)
	tp.On("Save",
		&model.Tournament{ID: 3, Participants: []*model.Player{{ID: 3}},
			Winner: model.Winner{Player: &model.Player{ID: 3}}}).Return(
		fmt.Errorf("failed"))
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
			expectedErrMsg: "tour 2: load: failed",
		},
		{
			name:           "save error",
			ID:             3,
			expectedErrMsg: "tour 3: save: failed",
		},
	}

	service := Service{tp: tp}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := service.EndTournament(tc.ID)
			if err != nil {
				assert.EqualError(t, err, tc.expectedErrMsg)
			}
		})
	}
}

// What I should test in New method?
func TestNew(t *testing.T) {
	pp := new(mocks.PlayerProvider)
	tp := new(mocks.TourProvider)

	s := New(Builder{PP: pp, TP: tp})

	assert.Equal(t, &Service{pp: pp, tp: tp}, s)
}
