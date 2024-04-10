package handler

import (
	"encoding/json"
	"game-store-api/database"
	"game-store-api/service"
	"net/http"

	"github.com/gorilla/mux"
)

type ReviewHandler struct {
	review *service.ReviewService
}

func NewReviewHandler(review *service.ReviewService) *ReviewHandler {
	return &ReviewHandler{review: review}
}

func (rh *ReviewHandler) GetReviews(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	reviews, err := rh.review.GetReviews()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonM , err := json.Marshal(reviews)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonM)
}

func (rh *ReviewHandler) GetReview(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	id := r.URL.Query().Get("id")

	review, err := rh.review.GetReview(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsoM, err := json.Marshal(review)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsoM)
}

func (rh *ReviewHandler) CreateReview(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	defer r.Body.Close()

	var review database.Review
	err := json.NewDecoder(r.Body).Decode(&review)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	reviews, err := rh.review.CreateReview(review)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonM, err := json.Marshal(reviews)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonM)
}

func (rh *ReviewHandler) UpdateReview(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	defer r.Body.Close()

	var review database.Review
	err := json.NewDecoder(r.Body).Decode(&review)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	editedReview, err := rh.review.UpdateReview(review)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonM, err := json.Marshal(editedReview)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonM)
}

func (rh *ReviewHandler) DeleteReview(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r.Method)

	id := r.URL.Query().Get("id")

	reviews, err := rh.review.DeleteReview(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonM, err := json.Marshal(reviews)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonM)
}

func (rh *ReviewHandler) Routes(subRouter *mux.Router) {
	subRouter.HandleFunc("/{id}", rh.GetReview).Methods(http.MethodGet, http.MethodOptions)
	subRouter.HandleFunc("", rh.CreateReview).Methods(http.MethodPost, http.MethodOptions)
	subRouter.HandleFunc("/{id}", rh.UpdateReview).Methods(http.MethodPut, http.MethodOptions)
	subRouter.HandleFunc("/{id}", rh.DeleteReview).Methods(http.MethodDelete, http.MethodOptions)
}