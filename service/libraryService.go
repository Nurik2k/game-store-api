package service

import (
	"game-store-api/database"
)

type LibraryService struct {
	library database.ILibrary
}

type ILibraryService interface {
	GetLibraryByUser(userID int) ([]database.Library, error)
	CreateLibrary(library database.Library) (string, error)
	AddGameToLibraryFromUser(library database.Library) (string, error)
	DeleteLibrary(id int) (string, error)
}

func NewLibraryService(library database.ILibrary) *LibraryService {
	return &LibraryService{library: library}
}

func (s *LibraryService) GetLibraryByUser(userID int) ([]database.Library, error) {
	libraries, err := s.library.GetLibrariesByUser(userID)
	if err != nil {
		return nil, err
	}

	return libraries, nil
}


func (s *LibraryService) CreateLibrary(library database.Library) (string, error) {
	err := s.library.CreateLibrary(library)
	if err != nil {
		return "", err
	}

	return "Library created", nil
}

func (s *LibraryService) AddGameToLibraryFromUser(library database.Library) (string, error) {
	err := s.library.AddGameToLibraryFromUser(library)
	if err != nil {
		return "", err
	}

	return "Game added to library", nil
}

func (s *LibraryService) DeleteLibrary(id int) (string, error) {
	err := s.library.DeleteLibrary(id)
	if err != nil {
		return "", err
	}

	return "Library deleted", nil
}