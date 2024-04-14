package handler

import (
	"encoding/json"
	"game-store-api/database"
	"game-store-api/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type FavoriteHandler struct {
	favorite *service.FavoriteService
}

func NewFavoriteHandler(favorite *service.FavoriteService) *FavoriteHandler {
	return &FavoriteHandler{favorite: favorite}
}

func (h *FavoriteHandler) GetFavoriteGamesByUser(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	favoriteGame, err := h.favorite.GetFavoriteGamesByUser(id)
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

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

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

func (h *FavoriteHandler) Routes(subRouter *mux.Router) {
	subRouter.HandleFunc("/{id}", h.GetFavoriteGamesByUser).Methods(http.MethodGet, http.MethodOptions)
	subRouter.HandleFunc("/{id}", h.AddFavorite).Methods(http.MethodPost, http.MethodOptions)
	subRouter.HandleFunc("/{id}", h.DeleteFavorite).Methods(http.MethodDelete, http.MethodOptions)
}