package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"leckomio.dev/go/museum/data"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contenct-Type", "application/json")
	// api/exhibitions?id=42
	id := r.URL.Query()["id"] // query returns a map
	if id != nil {
		finalId, err := strconv.Atoi(id[0])
		if err == nil && finalId < len(data.GetAll()) {
			json.NewEncoder(w).Encode(data.GetAll()[finalId])
		} else {
			http.Error(w, "Invalid exhibition", http.StatusBadRequest)
		}
	} else {
		json.NewEncoder(w).Encode(data.GetAll())
	}
}
