package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/Abinet16/rss/internal/database"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	// Implementation for creating a user will go here
	type Parameters struct {
		Name	 string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := Parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt:time.Now().UTC(),
		UpdatedAt: time.Now().UTC() ,
		Name: params.Name,
	})
	if err != nil{
		responseWithError(w,400,fmt.Sprintf("Couldn't create user: %v", err))
		return
	}
		responseWithJSON(w, 200,databaseUserToUser(user))
}


func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
responseWithJSON(w,200,databaseUserToUser(user))
}

