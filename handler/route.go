// Package route provides router and method to set handler for initial paths.
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

	API struct {
		s TourService
		r *mux.Router
	}
)

func New(s TourService) *API {
	a := &API{
		s: s,
		r: mux.NewRouter(),
	}
	a.initRouter()

	return a
}

func (a API) Router() *mux.Router {
	return a.r
}

// Serve sets endpoint handler and restrict some paths with query values.
func (a API) initRouter() {
	p := a.r.
		Queries("playerID", "{id:"+intRegex+"}",
			"points", "{points:"+floatRegex+"}").
		Methods(http.MethodPut).
		Subrouter()
	p.HandleFunc("/take", a.Take)
	p.HandleFunc("/fund", a.Fund)

	t := a.r.
		Queries("tournamentID", "{id:"+intRegex+"}").
		Subrouter()
	t.HandleFunc("/announceTournament", a.AnnounceTournament).
		Queries("deposit", "{deposit:"+floatRegex+"}").
		Methods(http.MethodPut)
	t.HandleFunc("/joinTournament", a.JoinTournament).
		Queries("playerID", "{playerID:"+intRegex+"}").
		Methods(http.MethodPut)
	t.HandleFunc("/endTournament", a.EndTournament).
		Methods(http.MethodPut)
	t.HandleFunc("/resultTournament", a.ResultTournament).
		Methods(http.MethodGet)

	a.r.HandleFunc("/balance", a.Balance).
		Queries("playerID", "{id:"+intRegex+"}").
		Methods(http.MethodGet)
}
