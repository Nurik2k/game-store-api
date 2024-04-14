package service

import (
	"game-store-api/database"
)

type FavoriteService struct {
	favorite database.IFavoriteGame
}

type IFavoriteService interface {
	GetFavoriteGamesByUser(userID int) ([]database.Game, error)
	AddFavoriteGameToUser(favorite database.FavoriteGame) (string, error)
	DeleteFavoriteGameFromUser(id int) (string, error)
}

func NewFavoriteService(favorite database.IFavoriteGame) *FavoriteService {
	return &FavoriteService{favorite: favorite}
}

func (s *FavoriteService) GetFavoriteGamesByUser(userID int) ([]database.Game, error) {
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

func (s *FavoriteService) DeleteFavoriteGameFromUser(id int) (string, error) {
	err := s.favorite.DeleteFavoriteGameFromUser(id)
	if err != nil {
		return "", err
	}

	return "Favorite deleted", nil
}