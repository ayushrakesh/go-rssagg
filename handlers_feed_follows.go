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

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request) {

	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Auth error- %s", err))
		return
	}
	user, er := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if er != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't found user with APIKey- %v", er))
		return
	}
	userId := user.ID

	type pr struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	prm := pr{}

	decoder := json.NewDecoder(r.Body)
	errr := decoder.Decode(&prm)

	if errr != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing json- %v", errr))
		return
	}

	feedFollow, errrr := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    userId,
		FeedID:    prm.FeedID,
	})
	if errrr != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create feed_follow- %v", errrr))
		return
	}
	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feedFollow))
}
