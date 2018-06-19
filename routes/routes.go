// Package routes provides Router and method to set handlers for initial paths.
package routes

import (
	"github.com/gorilla/mux"
	"github.com/kat6123/tournament/handlers"
)

var Router = mux.NewRouter()

// Set sets endpoint handlers and restrict some paths with query values.
func Set() {
	p := Router.
		Queries("playerId", "", "points", "").
		Methods("PATCH").
		Subrouter()
	p.HandleFunc("/take", handlers.Take)
	p.HandleFunc("/fund", handlers.Fund)

	t := Router.
		Queries("tournamentId", "").
		Subrouter()
	t.HandleFunc("/announceTournament", handlers.AnnounceTournament).
		Queries("deposit", "").
		Methods("PATCH")
	t.HandleFunc("/joinTournament", handlers.JoinTournament).
		Queries("playerId", "")
	t.HandleFunc("/resultTournament", handlers.ResultTournament).
		Methods("GET")

	Router.HandleFunc("/balance", handlers.Balance).
		Queries("playerId", "").
		Methods("GET")
}
