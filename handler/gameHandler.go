package handler

import (
	"encoding/json"
	"game-store-api/database"
	"game-store-api/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type GameHandler struct {
	game *service.GameSerice
}

func NewGameHandler(game *service.GameSerice) *GameHandler {
	return &GameHandler{game: game}
}

func (g *GameHandler) GetGames(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	games, err := g.game.GetGames()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsoM, err := json.Marshal(games)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsoM)
}

func (g *GameHandler) GetGamesByGenre(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	params := mux.Vars(r)
	query := params["genre"]

	games, err := g.game.GetGamesByGenre(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsoM, err := json.Marshal(games)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsoM)
}

func (g *GameHandler) GetGame(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	game, err := g.game.GetGame(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonM, err := json.Marshal(game)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonM)
}


func (g *GameHandler) CreateGame(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	defer r.Body.Close()

	var game database.Game
	err := json.NewDecoder(r.Body).Decode(&game)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	games, err := g.game.CreateGame(game)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsoM, err := json.Marshal(games)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsoM)
}

func (g *GameHandler) UpdateGame(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	defer r.Body.Close()

	var game database.Game
	err := json.NewDecoder(r.Body).Decode(&game)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	games, err := g.game.UpdateGame(game)
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

func (g *GameHandler) DeleteGame(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	games, err := g.game.DeleteGame(id)
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

func (g *GameHandler) Routes(subRouter *mux.Router) {
	subRouter.HandleFunc("", g.GetGames).Methods(http.MethodGet, http.MethodOptions)
	subRouter.HandleFunc("/{id}", g.GetGame).Methods(http.MethodGet, http.MethodOptions)
	subRouter.HandleFunc("", g.GetGamesByGenre).Methods(http.MethodGet, http.MethodOptions)

	authRouter := subRouter.PathPrefix("/admin").Subrouter()
	authRouter.Use(RequireAuth)
	authRouter.HandleFunc("", g.CreateGame).Methods(http.MethodPost, http.MethodOptions)
	authRouter.HandleFunc("/{id}", g.UpdateGame).Methods(http.MethodPut, http.MethodOptions)
	authRouter.HandleFunc("/{id}", g.DeleteGame).Methods(http.MethodDelete, http.MethodOptions)
}

