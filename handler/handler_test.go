package handler

import (
	"net/http"
	"testing"

	"net/http/httptest"

	"github.com/kat6123/tournament/handler/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAPI_Take(t *testing.T) {
	tt := []struct {
		name           string
		playerID       string
		points         string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "take 200 points",
			playerID:       "1",
			points:         "200",
			expectedStatus: http.StatusOK,
			expectedBody:   "{}\n",
		},
	}

	ts := new(mocks.TourService)
	ts.On("Take", 1, 200).Return(nil)
	api := API{s: ts}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(
				http.MethodPut, "/take?playerID="+tc.playerID+"&points="+tc.points, nil)
			require.Nil(t, err, "create test has failed: %v", err)

			w := httptest.NewRecorder()

			api.Take(w, req)
			assert.Equal(t, tc.expectedStatus, w.Code)
			assert.Equal(t, tc.expectedBody, w.Body.String())
		})
	}
}
