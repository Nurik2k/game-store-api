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

	publishers, err := p.publisherService.GetPublishers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonM, err := json.Marshal(publishers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonM)
}

func (p *PublisherHandler) GetPublisher(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	id := r.URL.Query().Get("id")

	publisher, err := p.publisherService.GetPublisher(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonM, err := json.Marshal(publisher)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonM)
}

func (p *PublisherHandler) CreatePublisher(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	defer r.Body.Close()

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

	jsonM, err := json.Marshal(publishers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonM)
}

func (p *PublisherHandler) UpdatePublisher(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	defer r.Body.Close()

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

	jsonM, err := json.Marshal(publishers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonM)
}

func (p *PublisherHandler) DeletePublisher(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	id := r.URL.Query().Get("id")

	publishers, err := p.publisherService.DeletePublisher(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonM, errr := json.Marshal(publishers)
	if errr != nil {
		http.Error(w, errr.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonM)
}

func (p *PublisherHandler) Routes(subRouter *mux.Router) {
	subRouter.HandleFunc("", p.GetPublishers).Methods(http.MethodGet, http.MethodOptions)
	subRouter.HandleFunc("/{id}", p.GetPublisher).Methods(http.MethodGet, http.MethodOptions)
	subRouter.HandleFunc("", p.CreatePublisher).Methods(http.MethodPost, http.MethodOptions)
	subRouter.HandleFunc("/{id}", p.UpdatePublisher).Methods(http.MethodPut, http.MethodOptions)
	subRouter.HandleFunc("/{id}", p.DeletePublisher).Methods(http.MethodDelete, http.MethodOptions)
}