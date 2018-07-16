package handler

import (
	"context"
	"net/http"
)

type (
	intKey   string
	floatKey string

	middleware func(http.HandlerFunc) http.HandlerFunc
)

func intFromContext(ctx context.Context, param string) int {
	return ctx.Value(intKey(param)).(int)
}

func float32FromContext(ctx context.Context, param string) float32 {
	return ctx.Value(floatKey(param)).(float32)
}

func queryInt(params ...string) middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			for i := range params {
				val, err := getIntQueryParam(params[i], w, r)
				if err != nil {
					jsonError(w, err.Error(), http.StatusBadRequest)
					return
				}

				ctx = context.WithValue(ctx, intKey(params[i]), val)
			}

			f(w, r.WithContext(ctx))
		}
	}
}

func queryFloat32(params ...string) middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			for i := range params {
				val, err := getFloat32QueryParam(params[i], w, r)
				if err != nil {
					jsonError(w, err.Error(), http.StatusBadRequest)
					return
				}

				ctx = context.WithValue(ctx, floatKey(params[i]), val)
			}

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
