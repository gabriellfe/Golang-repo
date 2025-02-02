package handler

import (
	"encoding/json"
	"gabriellfe/config"
	"net/http"
)

func LiveHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("live")
}

func RefreshHandler(w http.ResponseWriter, r *http.Request) {
	config.Refresh()
	w.WriteHeader(http.StatusOK)
}
