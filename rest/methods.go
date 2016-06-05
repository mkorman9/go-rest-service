package rest

import (
	"encoding/json"
	"net/http"
)

func RespondJson(w http.ResponseWriter, status int, entity interface{}) {
	w.WriteHeader(status)
	if entity != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		json.NewEncoder(w).Encode(entity)
	}
}

func ReadJson(req *http.Request, entity interface{}) error {
	return json.NewDecoder(req.Body).Decode(entity)
}