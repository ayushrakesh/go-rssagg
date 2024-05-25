package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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

	respondWithJSON(w, 200, databaseUserToUser(user))
}
