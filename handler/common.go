package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func jsonError(w http.ResponseWriter, errMsg string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(message{errMsg}); err != nil {
		panic(fmt.Sprintf("encode error to json has failed: %v", err))
	}
}

func jsonResponse(w http.ResponseWriter, v interface{}) {
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		jsonError(w, fmt.Sprintf("encode %s as json: %v", v, err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func getIntQueryParam(param string, w http.ResponseWriter, r *http.Request) (int, error) {
	p, err := strconv.Atoi(r.URL.Query().Get(param))
	if err != nil {
		err = fmt.Errorf("parse %q as int has failed: %v", param, err)
		jsonError(w, err.Error(), http.StatusBadRequest)
		return 0, err
	}

	return p, nil
}

// Before I have used mux.Vars to get query params, but error with context after tests had to use r.URL.Query()? What is context
func getFloat32QueryParam(param string, w http.ResponseWriter, r *http.Request) (float32, error) {
	p, err := strconv.ParseFloat(r.URL.Query().Get(param), 32)
	if err != nil {
		err = fmt.Errorf("parse %q as float32 has failed: %v", param, err)
		jsonError(w, err.Error(), http.StatusBadRequest)
		return 0, err
	}

	return float32(p), nil
}
