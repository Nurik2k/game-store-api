package service

import (
	"game-store-api/database"
)

type GameSerice struct {
	game database.IGame
}

type IGameService interface {
	GetGames() ([]database.Game, error)
	GetGame(id string) (database.Game, error)
	CreateGame(game database.Game) (database.Game, error)
	UpdateGame(game database.Game) (string, error)
	DeleteGame(id string) (string, error)
}

func NewGameService(game database.IGame) *GameSerice {
	return &GameSerice{game: game}
}

func (s *GameSerice) GetGames() ([]database.Game, error) {
	games, err := s.game.GetGames()
	if err != nil {
		return nil, err
	}

	return games, nil
}

func (s *GameSerice) GetGame(id string) (database.Game, error) {
	game, err := s.game.GetGame(id)
	if err != nil {
		return game, err
	}
	
	return game, nil
}

func (s *GameSerice) CreateGame(game database.Game) (database.Game, error) {
	err := s.game.CreateGame(game)
	if err != nil {
		return game, err
	}
	
	game, err = s.game.GetGame(game.ID)
	if err != nil {
		return game, err
	}

	return game, nil
}
func (s *GameSerice) UpdateGame(game database.Game) (string, error) {
	err := s.game.UpdateGame(game)
	if err != nil {
		return "", err
	}

	return "game updated", nil
}

func (s *GameSerice) DeleteGame(id string) (string, error) {
	err := s.game.DeleteGame(id)
	if err != nil {
		return "", err
	}

	return "game deleted", nil
}
