// Package handlers provides handlers for initial urls.
package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func Take(w http.ResponseWriter, r *http.Request) {
	playerId, err := getIntQueryParam("playerId", w, r)
	if err != nil {
		return
	}

	points, err := getIntQueryParam("points", w, r)
	if err != nil {
		return
	}

	fmt.Fprintf(w, "%d %d", playerId, points)
}

func Fund(w http.ResponseWriter, r *http.Request) {
	playerId, err := getIntQueryParam("playerId", w, r)
	if err != nil {
		return
	}

	points, err := getIntQueryParam("points", w, r)
	if err != nil {
		return
	}

	fmt.Fprintf(w, "%d %d", playerId, points)
}

func AnnounceTournament(w http.ResponseWriter, r *http.Request) {
	tournamentId, err := getIntQueryParam("tournamentId", w, r)
	if err != nil {
		return
	}

	deposit, err := getIntQueryParam("deposit", w, r)
	if err != nil {
		return
	}

	// announce
	fmt.Fprintf(w, "%d %d", tournamentId, deposit)
}

func JoinTournament(w http.ResponseWriter, r *http.Request) {
	tournamentId, err := getIntQueryParam("tournamentId", w, r)
	if err != nil {
		return
	}

	playerId, err := getIntQueryParam("playerId", w, r)
	if err != nil {
		return
	}

	// join()
	fmt.Fprintf(w, "%d %d", tournamentId, playerId)
}

func ResultTournament(w http.ResponseWriter, r *http.Request) {
	tournamentId, err := getIntQueryParam("tournamentId", w, r)
	if err != nil {
		return
	}

	// winners tournament id
	fmt.Fprintf(w, "tournament: %d", tournamentId)
}

func Balance(w http.ResponseWriter, r *http.Request) {
	playerId, err := getIntQueryParam("playerId", w, r)
	if err != nil {
		return
	}

	// balance(playerId)
	fmt.Fprintf(w, "%d", playerId)
}

func getIntQueryParam(param string, w http.ResponseWriter, r *http.Request) (int, error) {
	p := r.URL.Query().Get(param)

	res, err := strconv.Atoi(p)
	if err != nil {
		errStr := fmt.Sprintf("parse %q has failed: %v", param, err)
		http.Error(w, errStr, http.StatusBadRequest)
		return 0, err
	}

	return res, nil
}
