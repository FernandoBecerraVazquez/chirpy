package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type validateChirpResponse struct {
	CleanedBody string `json:"cleaned_body"`
}

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	type requestBody struct {
		Body string `json:"body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := requestBody{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "Invalid request payload")
		return
	}

	if len(params.Body) > 140 {
		respondWithError(w, 400, "Chirp is too long")
		return
	}

	cleanedBody := cleanBody(params.Body)

	msg := validateChirpResponse{CleanedBody: cleanedBody}
	respondWithJSON(w, 200, msg)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	type Error struct {
		Message string `json:"error"`
	}

	payload := Error{Message: msg}
	respondWithJSON(w, code, payload)

}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte(`{"error": "Something went wrong"}`))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func cleanBody(body string) string {
	wordsList := strings.Split(body, " ")
	for i, word := range wordsList {
		lowercaseWord := strings.ToLower(word)
		switch lowercaseWord {
		case "kerfuffle":
			wordsList[i] = "****"
		case "sharbert":
			wordsList[i] = "****"
		case "fornax":
			wordsList[i] = "****"
		}
	}
	return strings.Join(wordsList, " ")
}
