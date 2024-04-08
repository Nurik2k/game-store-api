package service

import (
	"encoding/json"
	"game-store-api/database"
)

type GenreService struct {
	genre database.IGenre
}

type IGenreService interface {
	GetGenres() ([]byte, error)
	GetGenre(id string) ([]byte, error)
	CreateGenre(genre database.Genre) ([]byte, error)
	UpdateGenre(genre database.Genre) ([]byte, error)
	DeleteGenre(id string) ([]byte, error)
}

func NewGenreService(genre database.IGenre) *GenreService {
	return &GenreService{genre: genre}
}

func (s *GenreService) GetGenres() ([]byte, error) {
	genres, err := s.genre.GetGenres()
	if err != nil {
		return nil, err
	}

	jsonM, err := json.Marshal(genres)
	if err != nil {
		return nil, err
	}

	return jsonM, nil
}

func (s *GenreService) GetGenre(id string) ([]byte, error) {
	genre, err := s.genre.GetGenre(id)
	if err != nil {
		return nil, err
	}
	
	jsonM, err := json.Marshal(genre)
	if err != nil {
		return nil, err
	}

	return jsonM, nil
}

func (s *GenreService) CreateGenre(genre database.Genre) ([]byte, error) {
	err := s.genre.CreateGenre(genre)
	if err != nil {
		return nil, err
	}
	
	genre, err = s.genre.GetGenre(genre.ID)
	if err != nil {
		return nil, err
	}
		
	jsonM, err := json.Marshal(genre)
	if err != nil {
		return nil, err
	}

	return jsonM, nil
}
func (s *GenreService) UpdateGenre(genre database.Genre) ([]byte, error) {
	err := s.genre.UpdateGenre(genre)
	if err != nil {
		return nil, err
	}

	jsonM, err := json.Marshal(genre)
	if err != nil {
		return nil, err
	}

	return jsonM, nil
}

func (s *GenreService) DeleteGenre(id string) ([]byte, error) {
	err := s.genre.DeleteGenre(id)
	if err != nil {
		return nil, err
	}

	return []byte("Genre deleted"), nil
}