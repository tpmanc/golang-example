package helpers

import (
	"encoding/json"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	
	json.NewEncoder(w).Encode(payload)
}

func ResponseError(w http.ResponseWriter, payload interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	
	json.NewEncoder(w).Encode(payload)
}

func ResponseForbidden(w http.ResponseWriter, text string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusForbidden)
	
	response := map[string] string {
		"text": text,
	}
	json.NewEncoder(w).Encode(response)
}