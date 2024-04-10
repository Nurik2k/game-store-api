package handler

import (
	"encoding/json"
	"fmt"
	"game-store-api/database"
	"game-store-api/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type GenreHandler struct {
	genre *service.GenreService
}

func NewGenreHandler(genre *service.GenreService) *GenreHandler {
	return &GenreHandler{genre: genre}
}

func (g *GenreHandler) GetGenres(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	genres, err := g.genre.GetGenres()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonM, err := json.Marshal(genres)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonM)
}

func (g *GenreHandler) GetGenre(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	params := mux.Vars(r)
	


	fmt.Println(params)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	genre, err := g.genre.GetGenre(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonM, err := json.Marshal(genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonM)
}

func (g *GenreHandler) CreateGenre(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	defer r.Body.Close()

	var genre database.Genre
	err := json.NewDecoder(r.Body).Decode(&genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	genres, err := g.genre.CreateGenre(genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonM, err := json.Marshal(genres)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonM)
}

func (g *GenreHandler) UpdateGenre(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	defer r.Body.Close()

	err := r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var genre database.Genre
	genres, err := g.genre.UpdateGenre(genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonM, err := json.Marshal(genres)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonM)
}

func (g *GenreHandler) DeleteGenre(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	id := r.URL.Query().Get("id")

	genres, err := g.genre.DeleteGenre(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonM, err := json.Marshal(genres)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonM)
}

func (g *GenreHandler) Routes(subRouter *mux.Router) {
	subRouter.HandleFunc("", g.GetGenres).Methods(http.MethodGet, http.MethodOptions)
	subRouter.HandleFunc("/{id}", g.GetGenre).Methods(http.MethodGet, http.MethodOptions)
	subRouter.HandleFunc("", g.CreateGenre).Methods(http.MethodPost, http.MethodOptions)
	subRouter.HandleFunc("/{id}", g.UpdateGenre).Methods(http.MethodPut, http.MethodOptions)
	subRouter.HandleFunc("/{id}", g.DeleteGenre).Methods(http.MethodDelete, http.MethodOptions)
}