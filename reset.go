package main

import (
	"net/http"
)

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		respondWithError(w, 403, "Forbidden")
		return
	}

	err := cfg.dbQueries.ResetUsers(r.Context())
	if err != nil {
		respondWithError(w, 500, "Failed to reset users")
		return
	}

	cfg.fileserverHits.Store(0)

	w.WriteHeader(200)
}
