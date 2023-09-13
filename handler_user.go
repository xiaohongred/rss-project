package main

import (
	"encoding/json"
	"fmt"
	"github.com/xiaohongred/rss-project/internal/database"
	"net/http"
	"strconv"
	"time"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %+v", err))
		return
	}
	_, err = apiCfg.DB.CreateUsers(r.Context(), database.CreateUsersParams{
		Name:     params.Name,
		CreateAt: time.Now().UTC(),
		UpdateAt: time.Now().UTC(),
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error  insert user: %+v", err))
		return
	}
	id, _ := apiCfg.DB.LastInsertedUserID(r.Context())
	user, err := apiCfg.DB.GetUserInfo(r.Context(), int32(id))
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error get insert user: %+v", err))
		return
	}
	responseWithJSON(w, 200, struct {
		ID        string
		Name      string
		UpdatedAt string
		CreatedAt string
	}{
		ID:        strconv.Itoa(int(user.ID)),
		Name:      user.Name,
		UpdatedAt: user.UpdateAt.String(),
		CreatedAt: user.CreateAt.String(),
	})
}
