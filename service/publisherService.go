package service

import (
	"game-store-api/database"
)

type PublisherSerice struct {
	publisher database.IPublisher
}

type IPublisherService interface {
	GetPublishers() ([]database.Publisher, error)
	GetPublisher(id int) (database.Publisher, error)
	CreatePublisher(publisher database.Publisher) (string, error)
	UpdatePublisher(publisher database.Publisher) (string, error)
	DeletePublisher(id int) (string, error)
}

func NewPublisherService(publisher database.IPublisher) *PublisherSerice {
	return &PublisherSerice{publisher: publisher}
}

func (s *PublisherSerice) GetPublishers() ([]database.Publisher, error) {
	publishers, err := s.publisher.GetPublishers()
	if err != nil {
		return nil, err
	}

	return publishers, nil
}

func (s *PublisherSerice) GetPublisher(id int) (database.Publisher, error) {
	publisher, err := s.publisher.GetPublisher(id)
	if err != nil {
		return publisher, err
	}

	return publisher, nil
}

func (s *PublisherSerice) CreatePublisher(publisher database.Publisher) (string, error) {
	err := s.publisher.CreatePublisher(publisher)
	if err != nil {
		return "", err
	}

	return "Publisher added", nil
}

func (s *PublisherSerice) UpdatePublisher(publisher database.Publisher) (string, error) {
	err := s.publisher.UpdatePublisher(publisher)
	if err != nil {
		return "", err
	}

	return "Publisher updated", nil
}

func (s *PublisherSerice) DeletePublisher(id int) (string, error) {
	err := s.publisher.DeletePublisher(id)
	if err != nil {
		return "", err
	}

	return "Publisher deleted", nil
}