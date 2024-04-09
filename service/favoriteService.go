package service

import (
	"game-store-api/database"
)

type FavoriteService struct {
	favorite database.IFavoriteGame
}

type IFavoriteService interface {
	GetFavoriteGamesByUser(userID string) ([]database.FavoriteGame, error)
	AddFavoriteGameToUser(favorite database.FavoriteGame) (string, error)
	DeleteFavoriteGameFromUser(id string) (string, error)
}

func NewFavoriteService(favorite database.IFavoriteGame) *FavoriteService {
	return &FavoriteService{favorite: favorite}
}

func (s *FavoriteService) GetFavoriteGamesByUser(userID string) ([]database.FavoriteGame, error) {
	favoriteGame, err := s.favorite.GetFavoriteGamesByUser(userID)
	if err != nil {
		return nil, err
	}

	return favoriteGame, nil	
}

func (s *FavoriteService) AddFavoriteGameToUser(favorite database.FavoriteGame) (string, error) {
	err := s.favorite.AddFavoriteGameToUser(favorite)
	if err != nil {
		return "", err
	}

	return "Favorite added", nil
}

func (s *FavoriteService) DeleteFavoriteGameFromUser(id string) (string, error) {
	err := s.favorite.DeleteFavoriteGameFromUser(id)
	if err != nil {
		return "", err
	}

	return "Favorite deleted", nil
}