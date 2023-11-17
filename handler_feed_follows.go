package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/prabal01pathak/scratch/internal/database"
)

func (apiCfg *apiConfig) handlerFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	// b, err := io.ReadAll(r.Body)
	// if err != nil {
	// }
	// log.Printf("body is: %v", string(b))
	// return
	decoder := json.NewDecoder(r.Body)
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	params := parameters{}
	decodeErr := decoder.Decode(&params)
	if decodeErr != nil {
		respondWithError(w, 400, fmt.Sprintf("Please check the body data: %v", decodeErr))
		return
	}
	feedFollow, feedErr := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if feedErr != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't Feed the user: %v", feedErr))
		return
	}
	respondWithJson(w, 201, databaseFeedFollowToFeedFollow(feedFollow))
	return
}

func (apiCfg *apiConfig) handlerGetAllFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, feedErr := apiCfg.DB.GetAllFeedFollow(r.Context(), user.ID)
	if feedErr != nil {
		respondWithError(w, 400, fmt.Sprintf("Got the error while getting the feed: %v", feedErr))
	}
	respondWithJson(w, 200, databaseFeedFollowsToFeedFollows(feeds))
}
