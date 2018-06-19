// Package handlers provides handlers for initial urls.
package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"encoding/json"

	"github.com/kat6123/tournament/logic"
)

func Take(w http.ResponseWriter, r *http.Request) {
	playerId, err := getIntQueryParam("playerId", w, r)
	if err != nil {
		return
	}

	points, err := getFloat32QueryParam("points", w, r)
	if err != nil {
		return
	}

	player := logic.GetPlayer(playerId)
	player.Take(points)
}

func Fund(w http.ResponseWriter, r *http.Request) {
	playerId, err := getIntQueryParam("playerId", w, r)
	if err != nil {
		return
	}

	points, err := getFloat32QueryParam("points", w, r)
	if err != nil {
		return
	}

	player := logic.GetPlayer(playerId)
	player.Fund(points)
}

func AnnounceTournament(w http.ResponseWriter, r *http.Request) {
	tournamentId, err := getIntQueryParam("tournamentId", w, r)
	if err != nil {
		return
	}

	deposit, err := getFloat32QueryParam("deposit", w, r)
	if err != nil {
		return
	}

	tour := logic.GetTournament(tournamentId)
	fmt.Fprintf(w, "%s", tour)
	tour.SetDeposit(deposit)
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

	tournament := logic.GetTournament(tournamentId)
	player := logic.GetPlayer(playerId)
	tournament.Join(player)
}

func ResultTournament(w http.ResponseWriter, r *http.Request) {
	tournamentId, err := getIntQueryParam("tournamentId", w, r)
	if err != nil {
		return
	}

	tournament := logic.GetTournament(tournamentId)
	b, err := json.Marshal(tournament.Winners)
	if err != nil {
		errStr := fmt.Sprintf("encode winners %s as json has failed: %v", tournament.Winners, err)
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func Balance(w http.ResponseWriter, r *http.Request) {
	playerId, err := getIntQueryParam("playerId", w, r)
	if err != nil {
		return
	}

	player := logic.GetPlayer(playerId)
	b, err := json.Marshal(player)
	if err != nil {
		errStr := fmt.Sprintf("encode player %s as json has failed: %v", player, err)
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func getIntQueryParam(param string, w http.ResponseWriter, r *http.Request) (int, error) {
	p := r.URL.Query().Get(param)

	res, err := strconv.Atoi(p)
	if err != nil {
		errStr := fmt.Sprintf("parse %q as int has failed: %v", param, err)
		http.Error(w, errStr, http.StatusBadRequest)
		return 0, err
	}

	return res, nil
}

func getFloat32QueryParam(param string, w http.ResponseWriter, r *http.Request) (float32, error) {
	p := r.URL.Query().Get(param)

	res, err := strconv.ParseFloat(p, 32)
	if err != nil {
		errStr := fmt.Sprintf("parse %q as float32 has failed: %v", param, err)
		http.Error(w, errStr, http.StatusBadRequest)
		return 0, err
	}

	return float32(res), nil
}
