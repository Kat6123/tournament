// Package handlers provides handlers for initial urls.
package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/globalsign/mgo"
	"github.com/gorilla/mux"
	"github.com/kat6123/tournament/logic"
)

func jsonError(w http.ResponseWriter, message string, code int) {
	err := struct {
		message string
	}{
		message: message,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	// What if error when encode 'error' ?
	json.NewEncoder(w).Encode(err)
}

func jsonResponse(w http.ResponseWriter, v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		jsonError(w, fmt.Sprintf("encode %s as json: %v", v, err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func Take(w http.ResponseWriter, r *http.Request) {
	playerID, err := getIntQueryParam("playerID", w, r)
	if err != nil {
		return
	}

	points, err := getFloat32QueryParam("points", w, r)
	if err != nil {
		return
	}

	if err := logic.Take(playerID, points); err != nil {
		status := http.StatusInternalServerError

		// Bad dependency
		if err == mgo.ErrNotFound {
			status = http.StatusNotFound
		}
		jsonError(w, fmt.Sprintf("take points has failed: %v", err), status)
		return
	}
}

func Fund(w http.ResponseWriter, r *http.Request) {
	playerID, err := getIntQueryParam("playerID", w, r)
	if err != nil {
		return
	}

	points, err := getFloat32QueryParam("points", w, r)
	if err != nil {
		return
	}

	if err := logic.Fund(playerID, points); err != nil {
		status := http.StatusInternalServerError

		// Bad dependency
		if err == mgo.ErrNotFound {
			status = http.StatusNotFound
		}
		jsonError(w, fmt.Sprintf("fund points has failed: %v", err), status)
		return
	}
}

func AnnounceTournament(w http.ResponseWriter, r *http.Request) {
	tournamentID, err := getIntQueryParam("tournamentID", w, r)
	if err != nil {
		return
	}

	deposit, err := getFloat32QueryParam("points", w, r)
	if err != nil {
		return
	}

	if err := logic.AnnounceTournament(tournamentID, deposit); err != nil {
		status := http.StatusInternalServerError

		// Bad dependency
		if mgo.IsDup(err) {
			status = http.StatusBadRequest
		}
		jsonError(w, fmt.Sprintf("announce tournament has failed: %v", err), status)
		return
	}

}

func JoinTournament(w http.ResponseWriter, r *http.Request) {
	tournamentID, err := getIntQueryParam("tournamentID", w, r)
	if err != nil {
		return
	}

	playerID, err := getIntQueryParam("playerID", w, r)
	if err != nil {
		return
	}

	if err := logic.JoinTournament(tournamentID, playerID); err != nil {
		status := http.StatusInternalServerError

		// Bad dependency
		if err == mgo.ErrNotFound {
			status = http.StatusNotFound
		}
		jsonError(w, fmt.Sprintf("join to tournament id %d of player id %d has failed: %v",
			tournamentID, playerID, err), status)
		return
	}
}

func EndTournament(w http.ResponseWriter, r *http.Request) {
	tournamentID, err := getIntQueryParam("tournamentID", w, r)
	if err != nil {
		return
	}

	if err := logic.EndTournament(tournamentID); err != nil {
		status := http.StatusInternalServerError

		// Bad dependency
		if err == mgo.ErrNotFound {
			status = http.StatusNotFound
		}
		jsonError(w, fmt.Sprintf("end tournament id %d has failed: %v", tournamentID, err), status)
		return
	}
}

func ResultTournament(w http.ResponseWriter, r *http.Request) {
	tournamentID, err := getIntQueryParam("tournamentID", w, r)
	if err != nil {
		return
	}

	winner, err := logic.ResultTournament(tournamentID)
	if err != nil {
		status := http.StatusInternalServerError

		// Bad dependency
		if err == mgo.ErrNotFound {
			status = http.StatusNotFound
		}
		jsonError(w, fmt.Sprintf("get result of tournament id %d has failed: %v", tournamentID, err), status)
		return
	}

	jsonResponse(w, winner)
}

func Balance(w http.ResponseWriter, r *http.Request) {
	playerID, err := getIntQueryParam("playerID", w, r)
	if err != nil {
		return
	}

	b, err := logic.Balance(playerID)
	if err != nil {
		status := http.StatusInternalServerError

		// Bad dependency
		if err == mgo.ErrNotFound {
			status = http.StatusNotFound
		}
		jsonError(w, fmt.Sprintf("load balance has failed: %v", err), status)
		return
	}

	balance := struct {
		PlayerId int     `json:"playerId"`
		Balance  float32 `json:"balance"`
	}{
		PlayerId: playerID,
		Balance:  b,
	}
	jsonResponse(w, balance)
}

func getIntQueryParam(param string, w http.ResponseWriter, r *http.Request) (int, error) {
	p, err := strconv.Atoi(mux.Vars(r)[param])
	if err != nil {
		// What return as error?
		err := fmt.Errorf("parse %q as int has failed: %v", param, err)
		jsonError(w, err.Error(), http.StatusBadRequest)
		return 0, err
	}

	return p, nil
}

func getFloat32QueryParam(param string, w http.ResponseWriter, r *http.Request) (float32, error) {
	p, err := strconv.ParseFloat(mux.Vars(r)[param], 32)
	if err != nil {
		err := fmt.Errorf("parse %q as float32 has failed: %v", param, err)
		jsonError(w, err.Error(), http.StatusBadRequest)
		return 0, err
	}

	return float32(p), nil
}
