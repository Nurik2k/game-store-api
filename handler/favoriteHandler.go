package handler

import (
	"encoding/json"
	"game-store-api/database"
	"game-store-api/service"
	"net/http"

	"github.com/gorilla/mux"
)

type FavoriteHandler struct {
	favorite *service.FavoriteService
}

func NewFavoriteHandler(favorite *service.FavoriteService) *FavoriteHandler {
	return &FavoriteHandler{favorite: favorite}
}

func (h *FavoriteHandler) GetFavorites(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path != "/api/v1/favorite-games/{user_id}" {
		http.Error(w, "404 Page not found!", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "This request not GET!", http.StatusMethodNotAllowed)
		return
	}

	userID := r.URL.Query().Get("user_id")

	games, err := h.favorite.GetFavoriteGamesByUser(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(games)
}

func (h *FavoriteHandler) AddFavorite(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	defer r.Body.Close()

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path != "/api/v1/add-favorite" {
		http.Error(w, "404 Page not found!", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "This request not POST!", http.StatusMethodNotAllowed)
		return
	}

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
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(games)
}

func (h *FavoriteHandler) DeleteFavorite(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path != "/api/v1/delete-favorite" {
		http.Error(w, "404 Page not found!", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodDelete {
		http.Error(w, "This request not DELETE!", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	games, err := h.favorite.DeleteFavoriteGameFromUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(games)
}

func (h *FavoriteHandler) favoriteRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/favorite-games/{user_id}", h.GetFavorites).Methods("GET")
	router.HandleFunc("/api/v1/add-favorite", h.AddFavorite).Methods("POST")
	router.HandleFunc("/api/v1/delete-favorite", h.DeleteFavorite).Methods("DELETE")

	return router
}