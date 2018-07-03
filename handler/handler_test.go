package handler

import (
	"net/http"
	"testing"

	"net/http/httptest"

	"io/ioutil"

	"fmt"

	"github.com/globalsign/mgo"
	"github.com/kat6123/tournament/handler/mocks"
	"github.com/kat6123/tournament/model"
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
			expectedBody:   "",
		},
		{
			name:           "wrong playerID param",
			playerID:       "abra",
			points:         "200",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "{\"message\":\"parse \\\"playerID\\\" as int has failed: strconv.Atoi: parsing \\\"abra\\\": invalid syntax\"}\n",
		},
		{
			name:           "wrong points param",
			playerID:       "1",
			points:         "abra",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "{\"message\":\"parse \\\"points\\\" as float32 has failed: strconv.ParseFloat: parsing \\\"abra\\\": invalid syntax\"}\n",
		},
		{
			name:     "take error",
			playerID: "2",
			points:   "200",
			// Is it an internal error? May be bad request?
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "{\"message\":\"take points has failed: insufficient funds\"}\n",
		},
		{
			name:           "not found error",
			playerID:       "353",
			points:         "200",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "{\"message\":\"take points has failed: not found\"}\n",
		},
	}

	ts := new(mocks.TourService)
	defer ts.AssertExpectations(t)
	ts.On("Take", 1, float32(200)).Return(nil)
	ts.On("Take", 2, float32(200)).Return(fmt.Errorf("insufficient funds"))
	ts.On("Take", 353, float32(200)).Return(mgo.ErrNotFound)

	api := API{s: ts}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(
				http.MethodPut, "/take?playerID="+tc.playerID+"&points="+tc.points, nil)
			require.NoError(t, err, "create test has failed: %v", err)

			w := httptest.NewRecorder()

			api.Take(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			require.NoError(t, err, "couldn't read response body: %v", err)

			assert.Equal(t, tc.expectedStatus, resp.StatusCode, "server error: %s", string(body))
			assert.Equal(t, tc.expectedBody, string(body))
		})
	}
}

func TestAPI_Fund(t *testing.T) {
	tt := []struct {
		name           string
		playerID       string
		points         string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "fund 200 points",
			playerID:       "1",
			points:         "200",
			expectedStatus: http.StatusOK,
			expectedBody:   "",
		},
		{
			name:           "wrong playerID param",
			playerID:       "abra",
			points:         "200",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "{\"message\":\"parse \\\"playerID\\\" as int has failed: strconv.Atoi: parsing \\\"abra\\\": invalid syntax\"}\n",
		},
		{
			name:           "wrong points param",
			playerID:       "1",
			points:         "abra",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "{\"message\":\"parse \\\"points\\\" as float32 has failed: strconv.ParseFloat: parsing \\\"abra\\\": invalid syntax\"}\n",
		},
		{
			name:     "fund error",
			playerID: "2",
			points:   "200",
			// Is it an internal error? May be bad request?
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "{\"message\":\"fund points has failed: fund error\"}\n",
		},
		{
			name:           "not found error",
			playerID:       "353",
			points:         "200",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "{\"message\":\"fund points has failed: not found\"}\n",
		},
	}

	ts := new(mocks.TourService)
	defer ts.AssertExpectations(t)
	ts.On("Fund", 1, float32(200)).Return(nil)
	ts.On("Fund", 2, float32(200)).Return(fmt.Errorf("fund error"))
	ts.On("Fund", 353, float32(200)).Return(mgo.ErrNotFound)

	api := API{s: ts}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(
				http.MethodPut, "/fund?playerID="+tc.playerID+"&points="+tc.points, nil)
			require.NoError(t, err, "create test has failed: %v", err)

			w := httptest.NewRecorder()

			api.Fund(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			require.NoError(t, err, "couldn't read response body: %v", err)

			assert.Equal(t, tc.expectedStatus, resp.StatusCode, "server error: %s", string(body))
			assert.Equal(t, tc.expectedBody, string(body))
		})
	}
}

func TestAPI_AnnounceTournament(t *testing.T) {
	tt := []struct {
		name           string
		tournamentID   string
		points         string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "announce tour with 200 deposit",
			tournamentID:   "1",
			points:         "200",
			expectedStatus: http.StatusOK,
			expectedBody:   "",
		},
		{
			name:           "wrong playerID param",
			tournamentID:   "abra",
			points:         "200",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "{\"message\":\"parse \\\"tournamentID\\\" as int has failed: strconv.Atoi: parsing \\\"abra\\\": invalid syntax\"}\n",
		},
		{
			name:           "wrong points param",
			tournamentID:   "1",
			points:         "abra",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "{\"message\":\"parse \\\"points\\\" as float32 has failed: strconv.ParseFloat: parsing \\\"abra\\\": invalid syntax\"}\n",
		},
		{
			name:           "announce error",
			tournamentID:   "2",
			points:         "200",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "{\"message\":\"announce tournament has failed: announce error\"}\n",
		},
		{
			name:           "dup error",
			tournamentID:   "353",
			points:         "200",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "{\"message\":\"announce tournament has failed: duplicate was found\"}\n",
		},
	}

	ts := new(mocks.TourService)
	defer ts.AssertExpectations(t)
	ts.On("AnnounceTournament", 1, float32(200)).Return(nil)
	ts.On("AnnounceTournament", 2, float32(200)).Return(fmt.Errorf("announce error"))
	ts.On("AnnounceTournament", 353, float32(200)).Return(
		&mgo.QueryError{Code: 11000, Message: "duplicate was found"})

	api := API{s: ts}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(
				http.MethodPut, "/announceTournament?tournamentID="+tc.tournamentID+"&points="+tc.points, nil)
			require.NoError(t, err, "create test has failed: %v", err)

			w := httptest.NewRecorder()

			api.AnnounceTournament(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			require.NoError(t, err, "couldn't read response body: %v", err)

			assert.Equal(t, tc.expectedStatus, resp.StatusCode, "server error: %s", string(body))
			assert.Equal(t, tc.expectedBody, string(body))
		})
	}
}

func TestAPI_JoinTournament(t *testing.T) {
	tt := []struct {
		name           string
		tournamentID   string
		playerID       string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "join player 2 to tour 1",
			tournamentID:   "1",
			playerID:       "2",
			expectedStatus: http.StatusOK,
			expectedBody:   "",
		},
		{
			name:           "wrong tourID param",
			tournamentID:   "abra",
			playerID:       "200",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "{\"message\":\"parse \\\"tournamentID\\\" as int has failed: strconv.Atoi: parsing \\\"abra\\\": invalid syntax\"}\n",
		},
		{
			name:           "wrong playerID param",
			tournamentID:   "1",
			playerID:       "abra",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "{\"message\":\"parse \\\"playerID\\\" as int has failed: strconv.Atoi: parsing \\\"abra\\\": invalid syntax\"}\n",
		},
		{
			name:           "join error",
			tournamentID:   "2",
			playerID:       "3",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "{\"message\":\"join to tournament id 2 of player id 3 has failed: join error\"}\n",
		},
		{
			name:           "not found error",
			tournamentID:   "3",
			playerID:       "4",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "{\"message\":\"join to tournament id 3 of player id 4 has failed: not found\"}\n",
		},
	}

	ts := new(mocks.TourService)
	defer ts.AssertExpectations(t)
	ts.On("JoinTournament", 1, 2).Return(nil)
	ts.On("JoinTournament", 2, 3).Return(fmt.Errorf("join error"))
	ts.On("JoinTournament", 3, 4).Return(mgo.ErrNotFound)

	api := API{s: ts}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(
				http.MethodPut, "/joinTournament?tournamentID="+tc.tournamentID+"&playerID="+tc.playerID, nil)
			require.NoError(t, err, "create test has failed: %v", err)

			w := httptest.NewRecorder()

			api.JoinTournament(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			require.NoError(t, err, "couldn't read response body: %v", err)

			assert.Equal(t, tc.expectedStatus, resp.StatusCode, "server error: %s", string(body))
			assert.Equal(t, tc.expectedBody, string(body))
		})
	}
}

func TestAPI_EndTournament(t *testing.T) {
	tt := []struct {
		name           string
		tournamentID   string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "end tour 1",
			tournamentID:   "1",
			expectedStatus: http.StatusOK,
			expectedBody:   "",
		},
		{
			name:           "wrong tourID param",
			tournamentID:   "abra",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "{\"message\":\"parse \\\"tournamentID\\\" as int has failed: strconv.Atoi: parsing \\\"abra\\\": invalid syntax\"}\n",
		},
		{
			name:           "end error",
			tournamentID:   "2",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "{\"message\":\"end tournament id 2 has failed: end error\"}\n",
		},
		{
			name:           "not found error",
			tournamentID:   "3",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "{\"message\":\"end tournament id 3 has failed: not found\"}\n",
		},
	}

	ts := new(mocks.TourService)
	defer ts.AssertExpectations(t)
	ts.On("EndTournament", 1).Return(nil)
	ts.On("EndTournament", 2).Return(fmt.Errorf("end error"))
	ts.On("EndTournament", 3).Return(mgo.ErrNotFound)

	api := API{s: ts}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(
				http.MethodPut, "/endTournament?tournamentID="+tc.tournamentID, nil)
			require.NoError(t, err, "create test has failed: %v", err)

			w := httptest.NewRecorder()

			api.EndTournament(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			require.NoError(t, err, "couldn't read response body: %v", err)

			assert.Equal(t, tc.expectedStatus, resp.StatusCode, "server error: %s", string(body))
			assert.Equal(t, tc.expectedBody, string(body))
		})
	}
}

func TestAPI_ResultTournament(t *testing.T) {
	tt := []struct {
		name           string
		tournamentID   string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "result of first tour",
			tournamentID:   "1",
			expectedStatus: http.StatusOK,
			expectedBody:   "{\"playerId\":1,\"balance\":500,\"Prize\":400}",
		},
		{
			name:           "wrong tourID param",
			tournamentID:   "abra",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "{\"message\":\"parse \\\"tournamentID\\\" as int has failed: strconv.Atoi: parsing \\\"abra\\\": invalid syntax\"}\n",
		},
		{
			name:           "result error",
			tournamentID:   "2",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "{\"message\":\"get result of tournament id 2 has failed: end error\"}\n",
		},
		{
			name:           "not found error",
			tournamentID:   "3",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "{\"message\":\"get result of tournament id 3 has failed: not found\"}\n",
		},
	}

	ts := new(mocks.TourService)
	defer ts.AssertExpectations(t)
	ts.On("ResultTournament", 1).Return(
		&model.Winner{
			Player: &model.Player{ID: 1, Balance: 500},
			Prize:  400,
		}, nil)
	ts.On("ResultTournament", 2).Return(
		nil, fmt.Errorf("end error"))
	ts.On("ResultTournament", 3).Return(
		nil, mgo.ErrNotFound)

	api := API{s: ts}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(
				http.MethodPut, "/resultTournament?tournamentID="+tc.tournamentID, nil)
			require.NoError(t, err, "create test has failed: %v", err)

			w := httptest.NewRecorder()

			api.ResultTournament(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			require.NoError(t, err, "couldn't read response body: %v", err)

			assert.Equal(t, tc.expectedStatus, resp.StatusCode, "server error: %s", string(body))
			assert.Equal(t, tc.expectedBody, string(body))
		})
	}
}

func TestAPI_Balance(t *testing.T) {
	tt := []struct {
		name           string
		playerID       string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "result of first tour",
			playerID:       "1",
			expectedStatus: http.StatusOK,
			expectedBody:   "{\"playerId\":1,\"balance\":300}",
		},
		{
			name:           "wrong tourID param",
			playerID:       "abra",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "{\"message\":\"parse \\\"playerID\\\" as int has failed: strconv.Atoi: parsing \\\"abra\\\": invalid syntax\"}\n",
		},
		{
			name:           "balance error",
			playerID:       "2",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "{\"message\":\"load balance has failed: balance error\"}\n",
		},
		{
			name:           "not found error",
			playerID:       "3",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "{\"message\":\"load balance has failed: not found\"}\n",
		},
	}

	ts := new(mocks.TourService)
	defer ts.AssertExpectations(t)
	ts.On("Balance", 1).Return(float32(300), nil)
	ts.On("Balance", 2).Return(float32(0), fmt.Errorf("balance error"))
	ts.On("Balance", 3).Return(float32(0), mgo.ErrNotFound)

	api := API{s: ts}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(
				http.MethodPut, "/balance?playerID="+tc.playerID, nil)
			require.NoError(t, err, "create test has failed: %v", err)

			w := httptest.NewRecorder()

			api.Balance(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			require.NoError(t, err, "couldn't read response body: %v", err)

			assert.Equal(t, tc.expectedStatus, resp.StatusCode, "server error: %s", string(body))
			assert.Equal(t, tc.expectedBody, string(body))
		})
	}
}
