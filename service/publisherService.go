package service

import (
	"encoding/json"
	"game-store-api/database"
)

type PublisherSerice struct {
	publisher database.IPublisher
}

type IPublisherService interface {
	GetPublishers() ([]byte, error)
	GetPublisher(id string) ([]byte, error)
	CreatePublisher(publisher database.Publisher) ([]byte, error)
	UpdatePublisher(publisher database.Publisher) ([]byte, error)
	DeletePublisher(id string) ([]byte, error)
}

func NewPublisherService(publisher database.IPublisher) *PublisherSerice {
	return &PublisherSerice{publisher: publisher}
}

func (s *PublisherSerice) GetPublishers() ([]byte, error) {
	publishers, err := s.publisher.GetPublishers()
	if err != nil {
		return nil, err
	}

	jsonM, err := json.Marshal(publishers)
	if err != nil {
		return nil, err
	}

	return jsonM, nil
}

func (s *PublisherSerice) GetPublisher(id string) ([]byte, error) {
	publisher, err := s.publisher.GetPublisher(id)
	if err != nil {
		return nil, err
	}
	
	jsonM, err := json.Marshal(publisher)
	if err != nil {
		return nil, err
	}

	return jsonM, nil
}

func (s *PublisherSerice) CreatePublisher(publisher database.Publisher) ([]byte, error) {
	err := s.publisher.CreatePublisher(publisher)
	if err != nil {
		return nil, err
	}

	return []byte("Publisher added"), nil
}

func (s *PublisherSerice) UpdatePublisher(publisher database.Publisher) ([]byte, error) {
	err := s.publisher.UpdatePublisher(publisher)
	if err != nil {
		return nil, err
	}

	jsonM, err := json.Marshal(publisher)
	if err != nil {
		return nil, err
	}

	return jsonM, nil
}

func (s *PublisherSerice) DeletePublisher(id string) ([]byte, error) {
	err := s.publisher.DeletePublisher(id)
	if err != nil {
		return nil, err
	}

	return []byte("Publisher deleted"), nil
}