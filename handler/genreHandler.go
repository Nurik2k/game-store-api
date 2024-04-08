package handler

import (
	"encoding/json"
	"game-store-api/database"
	"game-store-api/service"
	"net/http"

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

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path != "/api/v1/genres" {
		http.Error(w, "404 Page not found!", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "This request not GET!", http.StatusMethodNotAllowed)
		return
	}

	genres, err := g.genre.GetGenres()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(genres)
}

func (g *GenreHandler) GetGenre(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path != "/api/v1/genre/{id}" {
		http.Error(w, "404 Page not found!", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "This request not GET!", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	genre, err := g.genre.GetGenre(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(genre)
}

func (g *GenreHandler) CreateGenre(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	defer r.Body.Close()

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path != "/api/v1/create-genre" {
		http.Error(w, "404 Page not found!", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "This request not POST!", http.StatusMethodNotAllowed)
		return
	}

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
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(genres)
}

func (g *GenreHandler) UpdateGenre(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	defer r.Body.Close()

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path != "/api/v1/update-genre/{id}" {
		http.Error(w, "404 Page not found!", http.StatusNotFound)
		return
	}	

	if r.Method != http.MethodPut {
		http.Error(w, "This request not PUT!", http.StatusMethodNotAllowed)
		return
	}

	var genre database.Genre
	err := r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	genres, err := g.genre.UpdateGenre(genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(genres)
}

func (g *GenreHandler) DeleteGenre(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path != "/api/v1/delete-genre/{id}" {
		http.Error(w, "404 Page not found!", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodDelete {
		http.Error(w, "This request not DELETE!", http.StatusMethodNotAllowed)
		return
	}	

	id := r.URL.Query().Get("id")

	genres, err := g.genre.DeleteGenre(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(genres)
}

func (g *GenreHandler) genreRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/genres", g.GetGenres).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/genre/{id}", g.GetGenre).Methods("GET" , "OPTIONS")
	router.HandleFunc("/api/v1/create-genre", g.CreateGenre).Methods("POST" , "OPTIONS")
	router.HandleFunc("/api/v1/update-genre/{id}", g.UpdateGenre).Methods("PUT" , "OPTIONS")
	router.HandleFunc("/api/v1/delete-genre/{id}", g.DeleteGenre).Methods("DELETE" , "OPTIONS")

	return router
}