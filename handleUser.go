package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/prabal01pathak/scratch/internal/database"
)

func (apiCfg *apiConfig) handlerUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	type parameters struct {
		Name string `json:"name"`
	}
	params := parameters{}
	decodeErr := decoder.Decode(&params)
	if decodeErr != nil {
		respondWithError(w, 400, fmt.Sprintf("Please check the body data: %v", decodeErr))
		return
	}
	user, userErr := apiCfg.DB.CreateUsers(r.Context(), database.CreateUsersParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if userErr != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create the user: %v", userErr))
		return
	}
	respondWithJson(w, 201, databaseUsertoUser(user))
	return
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJson(w, 200, databaseUsertoUser(user))
}
