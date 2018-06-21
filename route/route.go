// Package route provides router and method to set handler for initial paths.
package route

import (
	"github.com/gorilla/mux"
	"github.com/kat6123/tournament/handler"
)

const (
	intRegex   = "[0-9]+"
	floatRegex = "(?:[0-9]*[.])?[0-9]+"
)

// Serve sets endpoint handler and restrict some paths with query values.
func Serve() *mux.Router {
	router := mux.NewRouter()

	p := router.
		Queries("playerId", "{id:"+intRegex+"}",
			"points", "{points:"+floatRegex+"}").
		Methods("PUT").
		Subrouter()
	p.HandleFunc("/take", handler.Take)
	p.HandleFunc("/fund", handler.Fund)

	t := router.
		Queries("tournamentId", "{id:"+intRegex+"}").
		Subrouter()
	t.HandleFunc("/announceTournament", handler.AnnounceTournament).
		Queries("deposit", "{deposit:"+floatRegex+"}").
		Methods("PUT")
	t.HandleFunc("/joinTournament", handler.JoinTournament).
		Queries("playerId", "{playerId:"+intRegex+"}").
		Methods("PUT")
	t.HandleFunc("/endTournament", handler.EndTournament).
		Methods("PUT")
	t.HandleFunc("/resultTournament", handler.ResultTournament).
		Methods("GET")

	router.HandleFunc("/balance", handler.Balance).
		Queries("playerId", "{id:"+intRegex+"}").
		Methods("GET")

	return router
}
