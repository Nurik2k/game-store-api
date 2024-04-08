package service

import (
	"encoding/json"
	"game-store-api/database"
)

type ReviewService struct {
	review database.IReview
}

type IReviewService interface {
	GetReviews() ([]byte, error)
	GetReview(id string) ([]byte, error)
	CreateReview(review database.Review) ([]byte, error)
	UpdateReview(review database.Review) ([]byte, error)
	DeleteReview(id string) ([]byte, error)
}

func NewReviewService(review database.IReview) *ReviewService {
	return &ReviewService{review: review}
}

func (s *ReviewService) GetReviews() ([]byte, error) {
	reviews, err := s.review.GetReviews()
	if err != nil {
		return nil, err
	}

	jsonM, err := json.Marshal(reviews)
	if err != nil {
		return nil, err
	}

	return jsonM, nil
}

func (s *ReviewService) GetReview(id string) ([]byte, error) {
	review, err := s.review.GetReview(id)
	if err != nil {
		return nil, err
	}

	jsonM, err := json.Marshal(review)
	if err != nil {
		return nil, err
	}

	return jsonM, nil
}
func (s *ReviewService) CreateReview(review database.Review) ([]byte, error) {
	err := s.review.CreateReview(review)
	if err != nil {
		return nil, err
	}

	return []byte("Review created"), nil
}

func (s *ReviewService) UpdateReview(review database.Review) ([]byte, error) {
	err := s.review.UpdateReview(review)
	if err != nil {
		return nil, err
	}

	return []byte("Review updated"), nil
}

func (s *ReviewService) DeleteReview(id string) ([]byte, error) {
	err := s.review.DeleteReview(id)
	if err != nil {
		return nil, err
	}

	return []byte("Review deleted"), nil
}