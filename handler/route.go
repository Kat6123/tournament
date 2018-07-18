// Package handler provides router and method to set handler for initial paths.
package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kat6123/tournament/model"
)

const (
	intRegex   = "[0-9]+"
	floatRegex = "(?:[0-9]*[.])?[0-9]+"
)

type (
	// TourService will used by handlers.
	TourService interface {
		// Take loads player from repository, takes points and saves it.
		Take(playerID int, points float32) error
		// Fund loads player from repository, funds points and saves it.
		Fund(playerID int, points float32) error
		// AnnounceTournament creates tour and saves it in db. If already exists than return an error.
		AnnounceTournament(tourID int, deposit float32) error
		// JoinTournament joins player to tour and saves tour in db.
		JoinTournament(tourID, playerID int) error
		// Balance loads player and returns it balance.
		Balance(playerID int) (float32, error)
		// ResultTournament loads tour and returns its winner.
		ResultTournament(tourID int) (*model.Winner, error)
		// EndTournament loads tour, ends it and saves in db.
		EndTournament(tourID int) error
	}

	// API providesmethods to work with handler package.
	API struct {
		s TourService
		r *mux.Router
	}
)

// New returns new API instance and initializes api router.
func New(s TourService) *API {
	a := &API{
		s: s,
		r: mux.NewRouter(),
	}
	a.initRouter()

	return a
}

// Router returns Router of API instance. It will be initialized after New method is called.
// What better to return mux.Router or http.Handler
func (a API) Router() http.Handler {
	return a.r
}

// Serve sets endpoint handler and restrict some paths with query values.
func (a API) initRouter() {
	p := a.r.
		Queries("playerID", "{id:"+intRegex+"}",
			"points", "{points:"+floatRegex+"}").
		Methods(http.MethodPut).
		Subrouter()
	p.HandleFunc("/take", chain(a.Take, queryTake()))
	p.HandleFunc("/fund", chain(a.Fund, queryFund()))

	t := a.r.
		Queries("tournamentID", "{id:"+intRegex+"}").
		Subrouter()
	t.HandleFunc("/announceTournament",
		chain(a.AnnounceTournament, queryAnnounce())).
		Queries("deposit", "{deposit:"+floatRegex+"}").
		Methods(http.MethodPut)
	t.HandleFunc("/joinTournament",
		chain(a.JoinTournament, queryJoin())).
		Queries("playerID", "{playerID:"+intRegex+"}").
		Methods(http.MethodPut)
	t.HandleFunc("/endTournament",
		chain(a.EndTournament, queryEnd())).
		Methods(http.MethodPut)
	t.HandleFunc("/resultTournament",
		chain(a.ResultTournament, queryResult())).
		Methods(http.MethodGet)

	a.r.HandleFunc("/balance",
		chain(a.Balance, queryBalance())).
		Queries("playerID", "{id:"+intRegex+"}").
		Methods(http.MethodGet)
}
