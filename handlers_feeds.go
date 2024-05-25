package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ayushrakesh/go-rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request) {

	val := r.Header.Get("Authorization")
	if val == "" {
		respondWithError(w, 403, "Auth error")
		return
	}
	vals := strings.Split(val, " ")
	apiKey := vals[1]

	user, errr := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if errr != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't found user with APIKey- %v", errr))
		return
	}
	type params struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	pr := params{}

	decoder := json.NewDecoder(r.Body)

	er := decoder.Decode(&pr)
	if er != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing json- %v", er))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      pr.Name,
		Url:       pr.Url,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create feed- %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {

	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feeds- %s", err))
		return
	}

	respondWithJSON(w, 200, databaseFeedsToFeeds(feeds))
}
