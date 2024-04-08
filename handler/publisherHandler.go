package handler

import (
	"encoding/json"
	"game-store-api/database"
	"game-store-api/service"
	"net/http"

	"github.com/gorilla/mux"
)

type PublisherHandler struct {
	publisherService service.IPublisherService
}

func NewPublisherHandler(publisherService service.IPublisherService) *PublisherHandler {
	return &PublisherHandler{publisherService: publisherService}
}

func (p *PublisherHandler) GetPublishers(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path != "/api/v1/publishers" {	
		http.Error(w, "404 Page not found!", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "This request not GET!", http.StatusMethodNotAllowed)
		return
	}

	publishers, err := p.publisherService.GetPublishers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(publishers)
}

func (p *PublisherHandler) GetPublisher(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path != "/api/v1/publisher/{id}" {
		http.Error(w, "404 Page not found!", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "This request not GET!", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	publisher, err := p.publisherService.GetPublisher(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(publisher)
}

func (p *PublisherHandler) CreatePublisher(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	defer r.Body.Close()

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path != "/api/v1/create-publisher" {
		http.Error(w, "404 Page not found!", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "This request not POST!", http.StatusMethodNotAllowed)
		return
	}

	var publisher database.Publisher
	err := json.NewDecoder(r.Body).Decode(&publisher)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	publishers, err := p.publisherService.CreatePublisher(publisher)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(publishers)
}

func (p *PublisherHandler) UpdatePublisher(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	defer r.Body.Close()

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path != "/api/v1/update-publisher/{id}" {
		http.Error(w, "404 Page not found!", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodPut {
		http.Error(w, "This request not PUT!", http.StatusMethodNotAllowed)
		return
	}

	var publisher database.Publisher
	err := json.NewDecoder(r.Body).Decode(&publisher)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	publishers, err := p.publisherService.UpdatePublisher(publisher)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(publishers)
}

func (p *PublisherHandler) DeletePublisher(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path != "/api/v1/delete-publisher/{id}" {
		http.Error(w, "404 Page not found!", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodDelete {
		http.Error(w, "This request not DELETE!", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	publishers, err := p.publisherService.DeletePublisher(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(publishers)
}

func (p *PublisherHandler) publisherRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/publishers", p.GetPublishers).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/publisher/{id}", p.GetPublisher).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/create-publisher", p.CreatePublisher).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/update-publisher/{id}", p.UpdatePublisher).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/v1/delete-publisher/{id}", p.DeletePublisher).Methods("DELETE", "OPTIONS")

	return router
}