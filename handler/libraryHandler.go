package handler

import (
	"encoding/json"
	"game-store-api/database"
	"game-store-api/service"
	"net/http"

	"github.com/gorilla/mux"
)

type LibraryHandler struct {
	library *service.LibraryService
}

func NewLibraryHandler(library *service.LibraryService) *LibraryHandler {
	return &LibraryHandler{library: library}
}

func (l *LibraryHandler) GetLibrariesByUser(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path != "/api/v1/library/{user_id}" {
		http.Error(w, "404 Page not found!", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "This request not GET!", http.StatusMethodNotAllowed)
		return
	}

	userID := r.URL.Query().Get("user_id")

	games, err := l.library.GetLibraryByUser(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}	

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(games)
}

func (l *LibraryHandler) CreateLibrary(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	defer r.Body.Close()

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path != "/api/v1/create-library" {
		http.Error(w, "404 Page not found!", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "This request not POST!", http.StatusMethodNotAllowed)
		return
	}

	var library database.Library
	err := r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}	

	games, err := l.library.CreateLibrary(library)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)	
	w.Write(games)
}

func (l *LibraryHandler) AddGameToLibraryFromUser(w http.ResponseWriter, r *http.Request) {	
	enableCors(w, r.Method)

	defer r.Body.Close()

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path != "/api/v1/add-library" {
		http.Error(w, "404 Page not found!", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "This request not POST!", http.StatusMethodNotAllowed)
		return
	}

	var library database.Library
	err := json.NewDecoder(r.Body).Decode(&library)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	games, err := l.library.AddGameToLibraryFromUser(library)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(games)
}

func (l *LibraryHandler) DeleteLibrary(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path != "/api/v1/delete-library/{id}" {
		http.Error(w, "404 Page not found!", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodDelete {
		http.Error(w, "This request not DELETE!", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	games, err := l.library.DeleteLibrary(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(games)
}

func (l *LibraryHandler) libraryRoutes()  *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/library/{user_id}", l.GetLibrariesByUser).Methods("GET")
	router.HandleFunc("/api/v1/create-library", l.CreateLibrary).Methods("POST")
	router.HandleFunc("/api/v1/add-library", l.AddGameToLibraryFromUser).Methods("POST")
	router.HandleFunc("/api/v1/delete-library/{id}", l.DeleteLibrary).Methods("DELETE")

	return router
}