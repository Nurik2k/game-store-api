package service

import (
	"encoding/json"
	"game-store-api/database"
)

type GameSerice struct {
	game database.IGame
}

type IGameService interface {
	GetGames() ([]byte, error)
	GetGame(id string) ([]byte, error)
	CreateGame(game database.Game) ([]byte, error)
	UpdateGame(game database.Game) ([]byte, error)
	DeleteGame(id string) ([]byte, error)
}

func NewGameService(game database.IGame) *GameSerice {
	return &GameSerice{game: game}
}

func (s *GameSerice) GetGames() ([]byte, error) {
	games, err := s.game.GetGames()
	if err != nil {
		return nil, err
	}

	jsonM, err := json.Marshal(games)
	if err != nil {
		return nil, err
	}

	return jsonM, nil
}

func (s *GameSerice) GetGame(id string) ([]byte, error) {
	game, err := s.game.GetGame(id)
	if err != nil {
		return nil, err
	}
	
	jsonM, err := json.Marshal(game)
	if err != nil {
		return nil, err
	}

	return jsonM, nil
}

func (s *GameSerice) CreateGame(game database.Game) ([]byte, error) {
	err := s.game.CreateGame(game)
	if err != nil {
		return nil, err
	}
	
	game, err = s.game.GetGame(game.ID)
	if err != nil {
		return nil, err
	}
		
	jsonM, err := json.Marshal(game)
	if err != nil {
		return nil, err
	}

	return jsonM, nil
}
func (s *GameSerice) UpdateGame(game database.Game) ([]byte, error) {
	err := s.game.UpdateGame(game)
	if err != nil {
		return nil, err
	}

	jsonM, err := json.Marshal(game)
	if err != nil {
		return nil, err
	}

	return jsonM, nil
}

func (s *GameSerice) DeleteGame(id string) ([]byte, error) {
	err := s.game.DeleteGame(id)
	if err != nil {
		return nil, err
	}

	return []byte("Game deleted"), nil
}
