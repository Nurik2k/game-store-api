package service

import (
	"game-store-api/database"
)

type GenreService struct {
	genre database.IGenre
}

type IGenreService interface {
	GetGenres() ([]database.Genre, error)
	GetGenre(id int) (database.Genre, error)
	CreateGenre(genre database.Genre) (database.Genre, error)
	UpdateGenre(genre database.Genre) (string, error)
	DeleteGenre(id int) (string, error)
}

func NewGenreService(genre database.IGenre) *GenreService {
	return &GenreService{genre: genre}
}

func (s *GenreService) GetGenres() ([]database.Genre, error) {
	genres, err := s.genre.GetGenres()
	if err != nil {
		return nil, err
	}

	return genres, nil
}

func (s *GenreService) GetGenre(id int) (database.Genre, error) {
	genre, err := s.genre.GetGenre(id)
	if err != nil {
		return genre, err
	}

	return genre, nil
}

func (s *GenreService) CreateGenre(genre database.Genre) (string, error) {
	msg,err := s.genre.CreateGenre(genre)
	if err != nil {
		return "", err
	}

	return msg, nil
}
func (s *GenreService) UpdateGenre(genre database.Genre) (string, error) {
	err := s.genre.UpdateGenre(genre)
	if err != nil {
		return 	"", err
	}

	return "Genre updated", nil
}

func (s *GenreService) DeleteGenre(id int) (string, error) {
	err := s.genre.DeleteGenre(id)
	if err != nil {
		return "", err
	}

	return "Genre deleted", nil
}