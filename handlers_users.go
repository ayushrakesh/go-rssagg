package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ayushrakesh/go-rssagg/internal/auth"
	"github.com/ayushrakesh/go-rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name string `json:"name"`
	}

	params := parameters{}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON - %s\n", err))
		return
	}

	user, er := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if er != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user %s\n", er))
		return
	}

	respondWithJSON(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) getUserByAPIKey(w http.ResponseWriter, r *http.Request) {

	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Authentication error- %v", err))
		return
	}

	user, er := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if er != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't found user- %v", er))
		return
	}
	respondWithJSON(w, 200, databaseUserToUser(user))
}
