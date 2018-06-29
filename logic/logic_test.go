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
			expectedErrMsg: "save ",
		},
	}

	pp := new(mocks.PlayerProvider)
	pp.On("ByID", 1).Return(&model.Player{ID: 1, Balance: 300}, nil)
	pp.On("Save", &model.Player{1, 300 - points}).Return(nil)

	pp.On("ByID", 2).Return(nil, fmt.Errorf("load failed"))

	// Balance less than points to cause take error
	pp.On("ByID", 3).Return(&model.Player{ID: 1, Balance: 100}, nil)

	pp.On("ByID", 4).Return(&model.Player{ID: 1, Balance: 300}, nil)
	pp.On("Save", &model.Player{1, 300 - points}).Return(
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
