package service

import (
	"encoding/json"
	"game-store-api/database"
)

type FavoriteService struct {
	favorite database.IFavoriteGame
}

type IFavoriteService interface {
	GetFavoriteGamesByUser(userID string) ([]byte, error)
	AddFavoriteGameToUser(favorite database.FavoriteGame) ([]byte, error)
	DeleteFavoriteGameFromUser(id string) ([]byte, error)
}

func NewFavoriteService(favorite database.IFavoriteGame) *FavoriteService {
	return &FavoriteService{favorite: favorite}
}

func (s *FavoriteService) GetFavoriteGamesByUser(userID string) ([]byte, error) {
	favoriteGame, err := s.favorite.GetFavoriteGamesByUser(userID)
	if err != nil {
		return nil, err
	}

	jsonM, err := json.Marshal(favoriteGame)
	if err != nil {
		return nil, err
	}

	return jsonM, nil
}

func (s *FavoriteService) AddFavoriteGameToUser(favorite database.FavoriteGame) ([]byte, error) {
	err := s.favorite.AddFavoriteGameToUser(favorite)
	if err != nil {
		return nil, err
	}

	return []byte("Favorite added"), nil
}

func (s *FavoriteService) DeleteFavoriteGameFromUser(id string) ([]byte, error) {
	err := s.favorite.DeleteFavoriteGameFromUser(id)
	if err != nil {
		return nil, err
	}

	return []byte("Favorite deleted"), nil
}