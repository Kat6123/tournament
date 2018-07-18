package handler

import (
	"context"
	"net/http"
)

type (
	takeKey     string
	fundKey     string
	announceKey string
	joinKey     string
	endKey      string
	resultKey   string
	balanceKey  string

	middleware func(http.HandlerFunc) http.HandlerFunc
)

func queryTake() middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			i, err := getIntQueryParam("playerID", w, r)
			if err != nil {
				jsonError(w, err.Error(), http.StatusBadRequest)
				return
			}
			ctx = context.WithValue(ctx, takeKey("playerID"), i)

			p, err := getFloat32QueryParam("points", w, r)
			if err != nil {
				jsonError(w, err.Error(), http.StatusBadRequest)
				return
			}
			ctx = context.WithValue(ctx, takeKey("points"), p)

			f(w, r.WithContext(ctx))
		}
	}
}

func queryFund() middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			i, err := getIntQueryParam("playerID", w, r)
			if err != nil {
				jsonError(w, err.Error(), http.StatusBadRequest)
				return
			}
			ctx = context.WithValue(ctx, fundKey("playerID"), i)

			p, err := getFloat32QueryParam("points", w, r)
			if err != nil {
				jsonError(w, err.Error(), http.StatusBadRequest)
				return
			}
			ctx = context.WithValue(ctx, fundKey("points"), p)

			f(w, r.WithContext(ctx))
		}
	}
}

func queryAnnounce() middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			i, err := getIntQueryParam("tournamentID", w, r)
			if err != nil {
				jsonError(w, err.Error(), http.StatusBadRequest)
				return
			}
			ctx = context.WithValue(ctx, announceKey("tournamentID"), i)

			p, err := getFloat32QueryParam("deposit", w, r)
			if err != nil {
				jsonError(w, err.Error(), http.StatusBadRequest)
				return
			}
			ctx = context.WithValue(ctx, announceKey("deposit"), p)

			f(w, r.WithContext(ctx))
		}
	}
}
func queryJoin() middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			i, err := getIntQueryParam("tournamentID", w, r)
			if err != nil {
				jsonError(w, err.Error(), http.StatusBadRequest)
				return
			}
			ctx = context.WithValue(ctx, joinKey("tournamentID"), i)

			i, err = getIntQueryParam("playerID", w, r)
			if err != nil {
				jsonError(w, err.Error(), http.StatusBadRequest)
				return
			}
			ctx = context.WithValue(ctx, joinKey("playerID"), i)

			f(w, r.WithContext(ctx))
		}
	}
}

func queryEnd() middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			i, err := getIntQueryParam("tournamentID", w, r)
			if err != nil {
				jsonError(w, err.Error(), http.StatusBadRequest)
				return
			}
			ctx = context.WithValue(ctx, endKey("tournamentID"), i)

			f(w, r.WithContext(ctx))
		}
	}
}

func queryResult() middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			i, err := getIntQueryParam("tournamentID", w, r)
			if err != nil {
				jsonError(w, err.Error(), http.StatusBadRequest)
				return
			}
			ctx = context.WithValue(ctx, resultKey("tournamentID"), i)

			f(w, r.WithContext(ctx))
		}
	}
}

func queryBalance() middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			i, err := getIntQueryParam("playerID", w, r)
			if err != nil {
				jsonError(w, err.Error(), http.StatusBadRequest)
				return
			}
			ctx = context.WithValue(ctx, balanceKey("playerID"), i)

			f(w, r.WithContext(ctx))
		}
	}
}

func chain(f http.HandlerFunc, mw ...middleware) http.HandlerFunc {
	for _, m := range mw {
		f = m(f)
	}

	return f
}
