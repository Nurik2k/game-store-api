package handler

import (
	"encoding/json"
	"game-store-api/database"
	"game-store-api/service"
	"net/http"
)

type FavoriteHandler struct {
	favorite *service.FavoriteService
}

func NewFavoriteHandler(favorite *service.FavoriteService) *FavoriteHandler {
	return &FavoriteHandler{favorite: favorite}
}

func (h *FavoriteHandler) GetFavorites(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	userID := r.URL.Query().Get("user_id")

	favoriteGame, err := h.favorite.GetFavoriteGamesByUser(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonM, err := json.Marshal(favoriteGame)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonM)
}

func (h *FavoriteHandler) AddFavorite(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	defer r.Body.Close()

	var favorite database.FavoriteGame
	err := json.NewDecoder(r.Body).Decode(&favorite)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	games, err := h.favorite.AddFavoriteGameToUser(favorite)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonM, err := json.Marshal(games)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonM)
}

func (h *FavoriteHandler) DeleteFavorite(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	id := r.URL.Query().Get("id")

	games, err := h.favorite.DeleteFavoriteGameFromUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonM, err := json.Marshal(games)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonM)
}
