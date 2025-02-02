package helper

import (
	"encoding/json"
	"net/http"
)

func Decode(r *http.Request, payload interface{}) error {
	return json.NewDecoder(r.Body).Decode(&payload)
}

func DecodeAndValidate(w http.ResponseWriter, r *http.Request, payload interface{}) error {
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return err
	}
	return nil
}

func EncodeStatusBody(w http.ResponseWriter, payload interface{}, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
