package service

import (
	"game-store-api/database"
)

type ReviewService struct {
	review database.IReview
}

type IReviewService interface {
	GetReviews() ([]database.Review, error)
	GetReview(id string) (database.Review, error)
	CreateReview(review database.Review) (string, error)
	UpdateReview(review database.Review) (string, error)
	DeleteReview(id string) (string, error)
}

func NewReviewService(review database.IReview) *ReviewService {
	return &ReviewService{review: review}
}

func (s *ReviewService) GetReviews() ([]database.Review, error) {
	reviews, err := s.review.GetReviews()
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (s *ReviewService) GetReview(id string) (database.Review, error) {
	review, err := s.review.GetReview(id)
	if err != nil {
		return review, err
	}

	return review, nil
}
func (s *ReviewService) CreateReview(review database.Review) (string, error) {
	err := s.review.CreateReview(review)
	if err != nil {
		return "", err
	}

	return "Review created", nil
}

func (s *ReviewService) UpdateReview(review database.Review) (string, error) {
	err := s.review.UpdateReview(review)
	if err != nil {
		return "", err
	}

	return "Review updated", nil
}

func (s *ReviewService) DeleteReview(id string) (string, error) {
	err := s.review.DeleteReview(id)
	if err != nil {
		return "", err
	}

	return "Review deleted", nil
}