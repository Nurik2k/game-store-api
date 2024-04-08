package service

import (
	"encoding/json"
	"game-store-api/database"
)

type LibraryService struct {
	library database.ILibrary
}

type ILibraryService interface {
	GetLibraryByUser(userID string) ([]byte, error)
	CreateLibrary(library database.Library) ([]byte, error)
	AddGameToLibraryFromUser(library database.Library) ([]byte, error)
	DeleteLibrary(id string) ([]byte, error)
}

func NewLibraryService(library database.ILibrary) *LibraryService {
	return &LibraryService{library: library}
}

func (s *LibraryService) GetLibraryByUser(userID string) ([]byte, error) {
	libraries, err := s.library.GetLibrariesByUser(userID)
	if err != nil {
		return nil, err
	}

	jsonM, err := json.Marshal(libraries)
	if err != nil {
		return nil, err
	}

	return jsonM, nil
}


func (s *LibraryService) CreateLibrary(library database.Library) ([]byte, error) {
	err := s.library.CreateLibrary(library)
	if err != nil {
		return nil, err
	}

	return []byte("Library created"), nil
}

func (s *LibraryService) AddGameToLibraryFromUser(library database.Library) ([]byte, error) {
	err := s.library.AddGameToLibraryFromUser(library)
	if err != nil {
		return nil, err
	}

	return []byte("Game added to library"), nil
}

func (s *LibraryService) DeleteLibrary(id string) ([]byte, error) {
	err := s.library.DeleteLibrary(id)
	if err != nil {
		return nil, err
	}

	return []byte("Library deleted"), nil
}