package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/prabal01pathak/scratch/internal/database"
)

func (apiCfg *apiConfig) handlerFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	decoder := json.NewDecoder(r.Body)
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	params := parameters{}
	decodeErr := decoder.Decode(&params)
	if decodeErr != nil {
		respondWithError(w, 400, fmt.Sprintf("Please check the body data: %v", decodeErr))
		return
	}
	feed, feedErr := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if feedErr != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't Feed the user: %v", feedErr))
		return
	}
	respondWithJson(w, 201, databaseFeedToFeed(feed))
	return
}

func (apiCfg *apiConfig) handleGetAllFeed(w http.ResponseWriter, r *http.Request) {
	feeds, feedErr := apiCfg.DB.GetAllFeed(r.Context())
	if feedErr != nil {
		respondWithError(w, 500, fmt.Sprintf("Got the error while getting the feed: %v", feedErr))
	}
	respondWithJson(w, 200, databasesFeedsToFeeds(feeds))
}
