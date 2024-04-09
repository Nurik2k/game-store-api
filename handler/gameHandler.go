package handler

import (
	"encoding/json"
	"game-store-api/database"
	"game-store-api/service"
	"net/http"

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

func (g *GameHandler) GetGame(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	id := r.URL.Query().Get("id")

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

	id := r.URL.Query().Get("id")

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

func (g *GameHandler) gameRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/games", g.GetGames).Methods("GET")
	router.HandleFunc("/api/v1/game/{id}", g.GetGame).Methods("GET")
	router.HandleFunc("/api/v1/create-game", g.CreateGame).Methods("POST")
	router.HandleFunc("/api/v1/update-game/{id}", g.UpdateGame).Methods("PUT")
	router.HandleFunc("/api/v1/delete-game/{id}", g.DeleteGame).Methods("DELETE")

	return router
}