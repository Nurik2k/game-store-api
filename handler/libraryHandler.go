package handler

import (
	"encoding/json"
	"game-store-api/database"
	"game-store-api/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type LibraryHandler struct {
	library *service.LibraryService
}

func NewLibraryHandler(library *service.LibraryService) *LibraryHandler {
	return &LibraryHandler{library: library}
}

func (l *LibraryHandler) GetLibraryByUser(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	games, err := l.library.GetLibraryByUser(id)
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

func (l *LibraryHandler) CreateLibrary(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	defer r.Body.Close()

	
	err := r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}	

	var library database.Library
	games, err := l.library.CreateLibrary(library)
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

func (l *LibraryHandler) AddGameToLibraryFromUser(w http.ResponseWriter, r *http.Request) {	
	enableCors(w, r.Method)

	defer r.Body.Close()

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

	jsonM, err := json.Marshal(games)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonM)
}

func (l *LibraryHandler) DeleteLibrary(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	games, err := l.library.DeleteLibrary(id)
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

func (l *LibraryHandler) Routes(subRouter *mux.Router) {
	subRouter.HandleFunc("/{id}", l.GetLibraryByUser).Methods(http.MethodGet, http.MethodOptions)
	subRouter.HandleFunc("", l.CreateLibrary).Methods(http.MethodPost, http.MethodOptions)
	subRouter.HandleFunc("/{id}", l.AddGameToLibraryFromUser).Methods(http.MethodPut, http.MethodOptions)
	subRouter.HandleFunc("/{id}", l.DeleteLibrary).Methods(http.MethodDelete, http.MethodOptions)
}