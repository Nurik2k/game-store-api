package handler

import (
	"encoding/json"
	"game-store-api/database"
	"game-store-api/service"
	"net/http"

	"github.com/gorilla/mux"
)

type PlatformHandler struct {
	platform *service.PlatformService
}

func NewPlatformHandler(platform *service.PlatformService) *PlatformHandler {
	return &PlatformHandler{platform: platform}
}

func (p *PlatformHandler) GetPlatforms(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	platforms, err := p.platform.GetPlatforms()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	jsonM, err := json.Marshal(platforms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonM)
}

func (p *PlatformHandler) GetPlatform(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	id := r.URL.Query().Get("id")

	platform, err := p.platform.GetPlatform(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonM, err := json.Marshal(platform)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonM)
}

func (p *PlatformHandler) CreatePlatform(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	defer r.Body.Close()

	var platform database.Platform
	err := json.NewDecoder(r.Body).Decode(&platform)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	platforms, err := p.platform.CreatePlatform(platform)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonM, err := json.Marshal(platforms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonM)
}

func (p *PlatformHandler) UpdatePlatform(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	defer r.Body.Close()

	var platform database.Platform
	err := r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	platforms, err := p.platform.UpdatePlatform(platform)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonM, err := json.Marshal(platforms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonM)
}

func (p *PlatformHandler) DeletePlatform(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	id := r.URL.Query().Get("id")

	platforms, err := p.platform.DeletePlatform(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonM, err := json.Marshal(platforms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonM)
}

func (p *PlatformHandler) Routes(subRouter *mux.Router) {
	subRouter.HandleFunc("", p.GetPlatforms).Methods(http.MethodGet, http.MethodOptions)
	subRouter.HandleFunc("/{id}", p.GetPlatform).Methods(http.MethodGet, http.MethodOptions)
	subRouter.HandleFunc("", p.CreatePlatform).Methods(http.MethodPost, http.MethodOptions)
	subRouter.HandleFunc("/{id}", p.UpdatePlatform).Methods(http.MethodPut, http.MethodOptions)
	subRouter.HandleFunc("/{id}", p.DeletePlatform).Methods(http.MethodDelete, http.MethodOptions)
}
