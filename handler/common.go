package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
