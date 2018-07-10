package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"io/ioutil"

	"github.com/gavv/httpexpect"
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
			expectedBody:   `{"message":"points was taken"}`,
		},
		{
			name:     "wrong playerID param",
			playerID: "abra",
			points:   "200",
			// Routing stops all bad requests for query string?
			expectedStatus: http.StatusNotFound,
			expectedBody:   `404 page not found`,
		},
		{
			name:           "wrong points param",
			playerID:       "1",
			points:         "abra",
			expectedStatus: http.StatusNotFound,
			expectedBody:   `404 page not found`,
		},
		{
			name:     "take error",
			playerID: "2",
			points:   "200",
			// Is it an internal error? May be bad request?
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"message":"take points has failed: insufficient funds"}`,
		},
		{
			name:           "not found error",
			playerID:       "353",
			points:         "200",
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"message":"take points has failed: not found"}`,
		},
	}

	ts := new(mocks.TourService)
	defer ts.AssertExpectations(t)
	ts.On("Take", 1, float32(200)).Return(nil)
	ts.On("Take", 2, float32(200)).Return(fmt.Errorf("insufficient funds"))
	ts.On("Take", 353, float32(200)).Return(mgo.ErrNotFound)

	server := httptest.NewServer(
		New(ts).Router())
	defer server.Close()

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			e := httpexpect.New(t, server.URL)

			resp := e.PUT("/take").
				WithQuery("playerID", tc.playerID).
				WithQuery("points", tc.points).Expect()

			resp.Status(tc.expectedStatus).Body().Equal(tc.expectedBody + "\n")
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
			expectedBody:   `{"message":"points was funded"}`,
		},
		{
			name:     "fund error",
			playerID: "2",
			points:   "200",
			// Is it an internal error? May be bad request?
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"message":"fund points has failed: fund error"}`,
		},
		{
			name:           "not found error",
			playerID:       "353",
			points:         "200",
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"message":"fund points has failed: not found"}`,
		},
	}

	ts := new(mocks.TourService)
	defer ts.AssertExpectations(t)
	ts.On("Fund", 1, float32(200)).Return(nil)
	ts.On("Fund", 2, float32(200)).Return(fmt.Errorf("fund error"))
	ts.On("Fund", 353, float32(200)).Return(mgo.ErrNotFound)

	server := httptest.NewServer(
		New(ts).Router())
	defer server.Close()

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			e := httpexpect.New(t, server.URL)

			resp := e.PUT("/fund").
				WithQuery("playerID", tc.playerID).
				WithQuery("points", tc.points).Expect()

			resp.Status(tc.expectedStatus).Body().Equal(tc.expectedBody + "\n")
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
			expectedBody:   `{"message":"tour was announced"}`,
		},
		{
			name:           "wrong tourID param",
			tournamentID:   "abra",
			points:         "200",
			expectedStatus: http.StatusNotFound,
			expectedBody:   `404 page not found`,
		},
		{
			name:           "wrong points param",
			tournamentID:   "1",
			points:         "abra",
			expectedStatus: http.StatusNotFound,
			expectedBody:   `404 page not found`,
		},
		{
			name:           "announce error",
			tournamentID:   "2",
			points:         "200",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"message":"announce tournament has failed: announce error"}`,
		},
		{
			name:           "dup error",
			tournamentID:   "353",
			points:         "200",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"message":"announce tournament has failed: duplicate was found"}`,
		},
	}

	ts := new(mocks.TourService)
	defer ts.AssertExpectations(t)
	ts.On("AnnounceTournament", 1, float32(200)).Return(nil)
	ts.On("AnnounceTournament", 2, float32(200)).Return(fmt.Errorf("announce error"))
	ts.On("AnnounceTournament", 353, float32(200)).Return(
		&mgo.QueryError{Code: 11000, Message: "duplicate was found"})

	server := httptest.NewServer(
		New(ts).Router())
	defer server.Close()

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			e := httpexpect.New(t, server.URL)

			resp := e.PUT("/announceTournament").
				WithQuery("tournamentID", tc.tournamentID).
				WithQuery("deposit", tc.points).Expect()

			resp.Status(tc.expectedStatus).Body().Equal(tc.expectedBody + "\n")
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
			expectedBody:   `{"message":"player was joined"}`,
		},
		{
			name:           "wrong tourID param",
			tournamentID:   "abra",
			playerID:       "200",
			expectedStatus: http.StatusNotFound,
			expectedBody:   `404 page not found`,
		},
		{
			name:           "wrong playerID param",
			tournamentID:   "1",
			playerID:       "abra",
			expectedStatus: http.StatusNotFound,
			expectedBody:   `404 page not found`,
		},
		{
			name:           "join error",
			tournamentID:   "2",
			playerID:       "3",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"message":"join to tournament id 2 of player id 3 has failed: join error"}`,
		},
		{
			name:           "not found error",
			tournamentID:   "3",
			playerID:       "4",
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"message":"join to tournament id 3 of player id 4 has failed: not found"}`,
		},
	}

	ts := new(mocks.TourService)
	defer ts.AssertExpectations(t)
	ts.On("JoinTournament", 1, 2).Return(nil)
	ts.On("JoinTournament", 2, 3).Return(fmt.Errorf("join error"))
	ts.On("JoinTournament", 3, 4).Return(mgo.ErrNotFound)

	server := httptest.NewServer(
		New(ts).Router())
	defer server.Close()

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			e := httpexpect.New(t, server.URL)

			resp := e.PUT("/joinTournament").
				WithQuery("tournamentID", tc.tournamentID).
				WithQuery("playerID", tc.playerID).Expect()

			resp.Status(tc.expectedStatus).Body().Equal(tc.expectedBody + "\n")
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
			expectedBody:   `{"message":"tour was ended"}`,
		},
		{
			name:           "wrong tourID param",
			tournamentID:   "abra",
			expectedStatus: http.StatusNotFound,
			expectedBody:   `404 page not found`,
		},
		{
			name:           "end error",
			tournamentID:   "2",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"message":"end tournament id 2 has failed: end error"}`,
		},
		{
			name:           "not found error",
			tournamentID:   "3",
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"message":"end tournament id 3 has failed: not found"}`,
		},
	}

	ts := new(mocks.TourService)
	defer ts.AssertExpectations(t)
	ts.On("EndTournament", 1).Return(nil)
	ts.On("EndTournament", 2).Return(fmt.Errorf("end error"))
	ts.On("EndTournament", 3).Return(mgo.ErrNotFound)

	server := httptest.NewServer(
		New(ts).Router())
	defer server.Close()

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			e := httpexpect.New(t, server.URL)

			resp := e.PUT("/endTournament").
				WithQuery("tournamentID", tc.tournamentID).Expect()

			resp.Status(tc.expectedStatus).Body().Equal(tc.expectedBody + "\n")
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
			expectedBody:   `{"playerId":1,"balance":500,"prize":400}`,
		},
		{
			name:           "wrong tourID param",
			tournamentID:   "abra",
			expectedStatus: http.StatusNotFound,
			expectedBody:   `404 page not found`,
		},
		{
			name:           "result error",
			tournamentID:   "2",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"message":"get result of tournament id 2 has failed: end error"}`,
		},
		{
			name:           "not found error",
			tournamentID:   "3",
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"message":"get result of tournament id 3 has failed: not found"}`,
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

	server := httptest.NewServer(
		New(ts).Router())
	defer server.Close()

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			e := httpexpect.New(t, server.URL)

			resp := e.GET("/resultTournament").
				WithQuery("tournamentID", tc.tournamentID).Expect()

			resp.Status(tc.expectedStatus).Body().Equal(tc.expectedBody + "\n")
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
			expectedBody:   `{"playerId":1,"balance":300}`,
		},
		{
			name:           "wrong tourID param",
			playerID:       "abra",
			expectedStatus: http.StatusNotFound,
			expectedBody:   `404 page not found`,
		},
		{
			name:           "balance error",
			playerID:       "2",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"message":"load balance has failed: balance error"}`,
		},
		{
			name:           "not found error",
			playerID:       "3",
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"message":"load balance has failed: not found"}`,
		},
	}

	ts := new(mocks.TourService)
	defer ts.AssertExpectations(t)
	ts.On("Balance", 1).Return(float32(300), nil)
	ts.On("Balance", 2).Return(float32(0), fmt.Errorf("balance error"))
	ts.On("Balance", 3).Return(float32(0), mgo.ErrNotFound)

	server := httptest.NewServer(
		New(ts).Router())
	defer server.Close()

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			e := httpexpect.New(t, server.URL)

			resp := e.GET("/balance").
				WithQuery("playerID", tc.playerID).Expect()

			resp.Status(tc.expectedStatus).Body().Equal(tc.expectedBody + "\n")
		})
	}
}

func Test_badQueryStrings(t *testing.T) {
	tt := []struct {
		name           string
		url            string
		qs             string
		expectedStatus int
	}{
		{
			name: "wrong playerID param in fund",
			url:  "/fund",
			qs:   "?playerID=abra&points=200",
			// Routing stops all bad requests for query string?
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "wrong points param in fund",
			qs:             "?playerID=1&points=abra",
			expectedStatus: http.StatusBadRequest,
		},
	}

	api := API{}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(
				http.MethodPut, tc.url+tc.qs, nil)
			require.NoError(t, err, "create test has failed: %v", err)

			w := httptest.NewRecorder()

			api.Fund(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			require.NoError(t, err, "couldn't read response body: %v", err)

			assert.Equal(t, tc.expectedStatus, resp.StatusCode, "server error: %s", string(body))
		})
	}
}
