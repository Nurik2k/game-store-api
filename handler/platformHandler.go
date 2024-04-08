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

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path != "/api/v1/platforms" {
		http.Error(w, "404 Page not found!", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "This request not GET!", http.StatusMethodNotAllowed)
		return
	}

	platforms, err := p.platform.GetPlatforms()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(platforms)
}

func (p *PlatformHandler) GetPlatform(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path != "/api/v1/platform/{id}" {
		http.Error(w, "404 Page not found!", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "This request not GET!", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	platform, err := p.platform.GetPlatform(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(platform)
}

func (p *PlatformHandler) CreatePlatform(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	defer r.Body.Close()

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path != "/api/v1/create-platform" {
		http.Error(w, "404 Page not found!", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "This request not POST!", http.StatusMethodNotAllowed)
		return
	}

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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(platforms)
}

func (p *PlatformHandler) UpdatePlatform(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	defer r.Body.Close()

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path != "/api/v1/update-platform/{id}" {
		http.Error(w, "404 Page not found!", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodPut {
		http.Error(w, "This request not PUT!", http.StatusMethodNotAllowed)
		return
	}

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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(platforms)
}

func (p *PlatformHandler) DeletePlatform(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path != "/api/v1/delete-platform/{id}" {
		http.Error(w, "404 Page not found!", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodDelete {
		http.Error(w, "This request not DELETE!", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	platforms, err := p.platform.DeletePlatform(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(platforms)
}

func (p *PlatformHandler) platformRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/platforms", p.GetPlatforms).Methods("GET")
	router.HandleFunc("/api/v1/platform/{id}", p.GetPlatform).Methods("GET")
	router.HandleFunc("/api/v1/create-platform", p.CreatePlatform).Methods("POST")
	router.HandleFunc("/api/v1/update-platform/{id}", p.UpdatePlatform).Methods("PUT")
	router.HandleFunc("/api/v1/delete-platform/{id}", p.DeletePlatform).Methods("DELETE")

	return router
}