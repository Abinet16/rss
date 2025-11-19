package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Abinet16/rss/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)


func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type Parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
			}

	var params Parameters
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		responseWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:      params.FeedID,
		UserID:    user.ID,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldn't create feed: %v", err))
		return
	}

	responseWithJSON(w, 201, databaseFeedFollowToFeedFollow(feedFollow))
}


func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiCfg.DB.GetFeedFollow(r.Context(), user.ID)
			if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldn't get feed follows: %v", err))
		return
	}



responseWithJSON(w, 201, databaseFeedFollowsToFeedFollows(feedFollows))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

feedFollowIDStr := chi.URLParam(r,"FeedFollowID")
feedFollowID, err := uuid.Parse(feedFollowIDStr)
if  err != nil {
	responseWithError(w, 400, fmt.Sprintf("Couldn't Parse feed Follow id: %v", err))
	return
}

err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
	ID: feedFollowID,
	UserID: user.ID,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldn't delete feed follow: %v",err))
		return
	}
	responseWithJSON(w, 200, struct{}{})
}