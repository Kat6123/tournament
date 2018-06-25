// Package route provides router and method to set handler for initial paths.
package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	intRegex   = "[0-9]+"
	floatRegex = "(?:[0-9]*[.])?[0-9]+"
)

// Serve sets endpoint handler and restrict some paths with query values.
func ServeRoutes() *mux.Router {
	router := mux.NewRouter()

	p := router.
		Queries("playerID", "{id:"+intRegex+"}",
			"points", "{points:"+floatRegex+"}").
		Methods(http.MethodPut).
		Subrouter()
	p.HandleFunc("/take", Take)
	p.HandleFunc("/fund", Fund)

	t := router.
		Queries("tournamentID", "{id:"+intRegex+"}").
		Subrouter()
	t.HandleFunc("/announceTournament", AnnounceTournament).
		Queries("deposit", "{deposit:"+floatRegex+"}").
		Methods(http.MethodPut)
	t.HandleFunc("/joinTournament", JoinTournament).
		Queries("playerID", "{playerID:"+intRegex+"}").
		Methods(http.MethodPut)
	t.HandleFunc("/endTournament", EndTournament).
		Methods(http.MethodPut)
	t.HandleFunc("/resultTournament", ResultTournament).
		Methods(http.MethodGet)

	router.HandleFunc("/balance", Balance).
		Queries("playerID", "{id:"+intRegex+"}").
		Methods(http.MethodGet)

	return router
}
