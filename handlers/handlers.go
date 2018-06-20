// Package handlers provides handlers for initial urls.
package handlers

import (
	"net/http"

	"fmt"
	"strconv"

	"encoding/json"

	"github.com/globalsign/mgo"
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

	if err := logic.Take(playerId, points); err != nil {
		status := http.StatusInternalServerError

		// Bad dependency
		if err == mgo.ErrNotFound {
			status = http.StatusNotFound
		}
		errStr := fmt.Sprintf("take points has failed: %v", err)
		http.Error(w, errStr, status)
		return
	}
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

	if err := logic.Fund(playerId, points); err != nil {
		status := http.StatusInternalServerError

		// Bad dependency
		if err == mgo.ErrNotFound {
			status = http.StatusNotFound
		}
		errStr := fmt.Sprintf("fund points has failed: %v", err)
		http.Error(w, errStr, status)
		return
	}
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

	if err := logic.AnnounceTournament(tournamentId, deposit); err != nil {
		status := http.StatusInternalServerError

		// Bad dependency
		if mgo.IsDup(err) {
			status = http.StatusBadRequest
		}
		errStr := fmt.Sprintf("announce tournament has failed: %v", err)
		http.Error(w, errStr, status)
		return
	}

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

	if err := logic.JoinTournament(tournamentId, playerId); err != nil {
		status := http.StatusInternalServerError

		// Bad dependency
		if err == mgo.ErrNotFound {
			status = http.StatusNotFound
		}
		errStr := fmt.Sprintf("join to tournament id %d of player id %d has failed: %v",
			tournamentId, playerId, err)
		http.Error(w, errStr, status)
		return
	}
}

func EndTournament(w http.ResponseWriter, r *http.Request) {
	tournamentId, err := getIntQueryParam("tournamentId", w, r)
	if err != nil {
		return
	}

	if err := logic.EndTournament(tournamentId); err != nil {
		status := http.StatusInternalServerError

		// Bad dependency
		if err == mgo.ErrNotFound {
			status = http.StatusNotFound
		}
		errStr := fmt.Sprintf("end tournament id %d has failed: %v",
			tournamentId, err)
		http.Error(w, errStr, status)
		return
	}
}

func ResultTournament(w http.ResponseWriter, r *http.Request) {
	tournamentId, err := getIntQueryParam("tournamentId", w, r)
	if err != nil {
		return
	}

	winner, err := logic.ResultTournament(tournamentId)
	if err != nil {
		status := http.StatusInternalServerError

		// Bad dependency
		if err == mgo.ErrNotFound {
			status = http.StatusNotFound
		}
		errStr := fmt.Sprintf("get result of tournament id %d has failed: %v",
			tournamentId, err)
		http.Error(w, errStr, status)
		return
	}

	b, err := json.Marshal(winner)
	if err != nil {
		errStr := fmt.Sprintf("encode winners %s as json has failed: %v", winner, err)
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

	// Better to call db.LoadPlayer in logic.Balance?
	player, err := logic.Balance(playerId)
	if err != nil {
		status := http.StatusInternalServerError

		// Bad dependency
		if err == mgo.ErrNotFound {
			status = http.StatusNotFound
		}
		errStr := fmt.Sprintf("load balance has failed: %v", err)
		http.Error(w, errStr, status)
		return
	}

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
