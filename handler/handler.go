// Package handler provides handlers for initial urls.
package handler

import (
	"fmt"
	"net/http"

	"github.com/kat6123/tournament/errors"
)

// Take handler takes player points.
func (a API) Take(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	playerID := ctx.Value(takeKey("playerID")).(int)
	points := ctx.Value(takeKey("points")).(float32)

	if err := a.s.Take(playerID, points); err != nil {
		jsonError(w, fmt.Sprintf("controller take: %v", err), errors.HTTPStatus(err))
		return
	}

	jsonResponse(w, message{"points was taken"})
}

// Fund handler funds points to player.
func (a API) Fund(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	playerID := ctx.Value(fundKey("playerID")).(int)
	points := ctx.Value(fundKey("points")).(float32)

	if err := a.s.Fund(playerID, points); err != nil {
		jsonError(w, fmt.Sprintf("controller fund: %v", err), errors.HTTPStatus(err))
		return
	}

	jsonResponse(w, message{"points was funded"})
}

// AnnounceTournament handler announces a new tournament.
func (a API) AnnounceTournament(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tournamentID := ctx.Value(announceKey("tournamentID")).(int)
	deposit := ctx.Value(announceKey("deposit")).(float32)

	if err := a.s.AnnounceTournament(tournamentID, deposit); err != nil {
		jsonError(w, fmt.Sprintf("controller announce: %v", err), errors.HTTPStatus(err))
		return
	}

	jsonResponse(w, message{"tour was announced"})
}

// JoinTournament handler joins player to tour.
func (a API) JoinTournament(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	playerID := ctx.Value(joinKey("playerID")).(int)
	tournamentID := ctx.Value(joinKey("tournamentID")).(int)

	if err := a.s.JoinTournament(tournamentID, playerID); err != nil {
		jsonError(w, fmt.Sprintf("controller join: %v", err), errors.HTTPStatus(err))
		return
	}

	jsonResponse(w, message{"player was joined"})
}

// EndTournament handler ends the tour.
func (a API) EndTournament(w http.ResponseWriter, r *http.Request) {
	tournamentID := r.Context().Value(endKey("tournamentID")).(int)

	if err := a.s.EndTournament(tournamentID); err != nil {
		jsonError(w, fmt.Sprintf("controller end: %v", err), errors.HTTPStatus(err))
		return
	}

	jsonResponse(w, message{"tour was ended"})
}

// ResultTournament handler returns the result of the tour.
func (a API) ResultTournament(w http.ResponseWriter, r *http.Request) {
	tournamentID := r.Context().Value(resultKey("tournamentID")).(int)

	winner, err := a.s.ResultTournament(tournamentID)
	if err != nil {
		jsonError(w, fmt.Sprintf("controller result: %v", err), errors.HTTPStatus(err))
		return
	}

	jsonResponse(w, winner)
}

// Balance handler returns the balance of the tour.
func (a API) Balance(w http.ResponseWriter, r *http.Request) {
	playerID := r.Context().Value(balanceKey("playerID")).(int)

	b, err := a.s.Balance(playerID)
	if err != nil {
		jsonError(w, fmt.Sprintf("controller balance: %v", err), errors.HTTPStatus(err))
		return
	}

	jsonResponse(w, balance{playerID, b})
}
