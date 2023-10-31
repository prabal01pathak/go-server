package main

import (
	"fmt"
	"net/http"

	"github.com/prabal01pathak/scratch/internal/auth"
	"github.com/prabal01pathak/scratch/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAuthHeader(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Bad Authentication: %v", err))
			return
		}
		fmt.Printf("api key is: %v", apiKey)
		user, userErr := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if userErr != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't get the user: %v", userErr))
			return
		}
		handler(w, r, user)
	}
}
