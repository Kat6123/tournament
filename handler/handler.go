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
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{message: %q}", message)
}

func jsonResponse(w http.ResponseWriter, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("encode %s as json: %v", v, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)

	return nil
}

func Take(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	playerId, err := strconv.Atoi(id)
	if err != nil {
		jsonError(w, fmt.Sprintf("parse %q as int has failed: %v", id, err), http.StatusBadRequest)
		return
	}

	p := mux.Vars(r)["points"]
	points, err := strconv.ParseFloat(p, 32)
	if err != nil {
		jsonError(w, fmt.Sprintf("parse %q as float32 has failed: %v", p, err), http.StatusBadRequest)
		return
	}

	if err := logic.Take(playerId, float32(points)); err != nil {
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
	id := mux.Vars(r)["id"]
	playerId, err := strconv.Atoi(id)
	if err != nil {
		jsonError(w, fmt.Sprintf("parse %q as int has failed: %v", id, err), http.StatusBadRequest)
		return
	}

	p := mux.Vars(r)["points"]
	points, err := strconv.ParseFloat(p, 32)
	if err != nil {
		jsonError(w, fmt.Sprintf("parse %q as float32 has failed: %v", p, err), http.StatusBadRequest)
		return
	}

	if err := logic.Fund(playerId, float32(points)); err != nil {
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
	id := mux.Vars(r)["id"]
	tournamentId, err := strconv.Atoi(id)
	if err != nil {
		jsonError(w, fmt.Sprintf("parse %q as int has failed: %v", id, err), http.StatusBadRequest)
		return
	}

	p := mux.Vars(r)["points"]
	deposit, err := strconv.ParseFloat(p, 32)
	if err != nil {
		jsonError(w, fmt.Sprintf("parse %q as float32 has failed: %v", p, err), http.StatusBadRequest)
		return
	}

	if err := logic.AnnounceTournament(tournamentId, float32(deposit)); err != nil {
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
	id := mux.Vars(r)["id"]
	tournamentId, err := strconv.Atoi(id)
	if err != nil {
		jsonError(w, fmt.Sprintf("parse %q as int has failed: %v", id, err), http.StatusBadRequest)
		return
	}

	pId := mux.Vars(r)["playerId"]
	playerId, err := strconv.Atoi(pId)
	if err != nil {
		jsonError(w, fmt.Sprintf("parse %q as int has failed: %v", id, err), http.StatusBadRequest)
		return
	}

	if err := logic.JoinTournament(tournamentId, playerId); err != nil {
		status := http.StatusInternalServerError

		// Bad dependency
		if err == mgo.ErrNotFound {
			status = http.StatusNotFound
		}
		jsonError(w, fmt.Sprintf("join to tournament id %d of player id %d has failed: %v",
			tournamentId, playerId, err), status)
		return
	}
}

func EndTournament(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	tournamentId, err := strconv.Atoi(id)
	if err != nil {
		jsonError(w, fmt.Sprintf("parse %q as int has failed: %v", id, err), http.StatusBadRequest)
		return
	}

	if err := logic.EndTournament(tournamentId); err != nil {
		status := http.StatusInternalServerError

		// Bad dependency
		if err == mgo.ErrNotFound {
			status = http.StatusNotFound
		}
		jsonError(w, fmt.Sprintf("end tournament id %d has failed: %v", tournamentId, err), status)
		return
	}
}

func ResultTournament(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	tournamentId, err := strconv.Atoi(id)
	if err != nil {
		jsonError(w, fmt.Sprintf("parse %q as int has failed: %v", id, err), http.StatusBadRequest)
		return
	}

	winner, err := logic.ResultTournament(tournamentId)
	if err != nil {
		status := http.StatusInternalServerError

		// Bad dependency
		if err == mgo.ErrNotFound {
			status = http.StatusNotFound
		}
		jsonError(w, fmt.Sprintf("get result of tournament id %d has failed: %v", tournamentId, err), status)
		return
	}

	if err := jsonResponse(w, winner); err != nil {
		jsonError(w, fmt.Sprintf("write winner %s to response has failed: %v", winner, err), http.StatusInternalServerError)
		return
	}
}

func Balance(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	playerId, err := strconv.Atoi(id)
	if err != nil {
		jsonError(w, fmt.Sprintf("parse %q as int has failed: %v", id, err), http.StatusBadRequest)
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
		jsonError(w, fmt.Sprintf("load balance has failed: %v", err), status)
		return
	}

	if err := jsonResponse(w, player); err != nil {
		jsonError(w, fmt.Sprintf("write player %s to response has failed: %v", player, err),
			http.StatusInternalServerError)
		return
	}
}
