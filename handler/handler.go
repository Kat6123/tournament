// Package handler provides handlers for initial urls.
package handler

import (
	"fmt"
	"net/http"

	"github.com/globalsign/mgo"
)

// Take handler takes player points.
func (a API) Take(w http.ResponseWriter, r *http.Request) {
	playerID := intFromContext(r.Context(), "playerID")
	points := float32FromContext(r.Context(), "points")

	if err := a.s.Take(playerID, points); err != nil {
		status := http.StatusInternalServerError

		// Bad dependency
		if err == mgo.ErrNotFound {
			status = http.StatusNotFound
		}
		jsonError(w, fmt.Sprintf("take points has failed: %v", err), status)
		return
	}

	jsonResponse(w, message{"points was taken"})
}

// Fund handler funds points to player.
func (a API) Fund(w http.ResponseWriter, r *http.Request) {
	playerID := intFromContext(r.Context(), "playerID")
	points := float32FromContext(r.Context(), "points")

	if err := a.s.Fund(playerID, points); err != nil {
		status := http.StatusInternalServerError

		// Bad dependency
		if err == mgo.ErrNotFound {
			status = http.StatusNotFound
		}
		jsonError(w, fmt.Sprintf("fund points has failed: %v", err), status)
		return
	}

	jsonResponse(w, message{"points was funded"})
}

// AnnounceTournament handler announces a new tournament.
func (a API) AnnounceTournament(w http.ResponseWriter, r *http.Request) {
	tournamentID := intFromContext(r.Context(), "tournamentID")
	deposit := float32FromContext(r.Context(), "deposit")

	if err := a.s.AnnounceTournament(tournamentID, deposit); err != nil {
		status := http.StatusInternalServerError

		// Bad dependency
		if mgo.IsDup(err) {
			status = http.StatusBadRequest
		}
		jsonError(w, fmt.Sprintf("announce tournament has failed: %v", err), status)
		return
	}

	jsonResponse(w, message{"tour was announced"})
}

// JoinTournament handler joins player to tour.
func (a API) JoinTournament(w http.ResponseWriter, r *http.Request) {
	playerID := intFromContext(r.Context(), "playerID")
	tournamentID := intFromContext(r.Context(), "tournamentID")

	if err := a.s.JoinTournament(tournamentID, playerID); err != nil {
		status := http.StatusInternalServerError

		// Bad dependency
		if err == mgo.ErrNotFound {
			status = http.StatusNotFound
		}
		jsonError(w, fmt.Sprintf("join to tournament id %d of player id %d has failed: %v",
			tournamentID, playerID, err), status)
		return
	}

	jsonResponse(w, message{"player was joined"})
}

// EndTournament handler ends the tour.
func (a API) EndTournament(w http.ResponseWriter, r *http.Request) {
	tournamentID := intFromContext(r.Context(), "tournamentID")

	if err := a.s.EndTournament(tournamentID); err != nil {
		status := http.StatusInternalServerError

		// Bad dependency
		if err == mgo.ErrNotFound {
			status = http.StatusNotFound
		}
		jsonError(w, fmt.Sprintf("end tournament id %d has failed: %v", tournamentID, err), status)
		return
	}

	jsonResponse(w, message{"tour was ended"})
}

// ResultTournament handler returns the result of the tour.
func (a API) ResultTournament(w http.ResponseWriter, r *http.Request) {
	tournamentID := intFromContext(r.Context(), "tournamentID")

	winner, err := a.s.ResultTournament(tournamentID)
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

// Balance handler returns the balance of the tour.
func (a API) Balance(w http.ResponseWriter, r *http.Request) {
	playerID := intFromContext(r.Context(), "playerID")

	b, err := a.s.Balance(playerID)
	if err != nil {
		status := http.StatusInternalServerError

		// Bad dependency
		if err == mgo.ErrNotFound {
			status = http.StatusNotFound
		}
		jsonError(w, fmt.Sprintf("load balance has failed: %v", err), status)
		return
	}

	jsonResponse(w, balance{playerID, b})
}
