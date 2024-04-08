package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type gameDb struct {
	db *pgxpool.Pool
}

type Game struct {
	ID          string       `db:"game_id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	GenreID     int       `db:"genre_id"`
	PublisherID int       `db:"publisher_id"`
	PlatformID  int       `db:"platform_id"`
	CreatedAt   time.Time `db:"created_at"`
}

type IGame interface {
	GetGames() ([]Game, error)
	GetGame(id string) (Game, error)
	CreateGame(game Game) error
	UpdateGame(game Game) error
	DeleteGame(id string) error
}

func (db *gameDb) GetGames() ([]Game, error) {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `SELECT game_id, name, description, genre_id, publisher_id, platform_id, created_at FROM games`
	rows, err := db.db.Query(ctx, queryRow)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []Game
	for rows.Next() {
		var game Game
		if err := rows.Scan(&game.ID, &game.Name, &game.Description, &game.GenreID, &game.PublisherID, &game.PlatformID, &game.CreatedAt); err != nil {
			return nil, err
		}
		games = append(games, game)
	}
	return games, nil
}

func (db *gameDb) GetGame(id string) (Game, error) {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `SELECT game_id, name, description, genre_id, publisher_id, platform_id, created_at FROM games WHERE game_id = $1`
	row := db.db.QueryRow(ctx, queryRow, id)
	var game Game
	if err := row.Scan(&game.ID, &game.Name, &game.Description, &game.GenreID, &game.PublisherID, &game.PlatformID, &game.CreatedAt); err != nil {
		return game, err
	}
	return game, nil
}

func (db *gameDb) CreateGame(game Game) error {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `INSERT INTO games (name, description, genre_id, publisher_id, platform_id) VALUES ($1, $2, $3, $4, $5)`
	_, err := db.db.Exec(ctx, queryRow, game.Name, game.Description, game.GenreID, game.PublisherID, game.PlatformID)
	return err
}

func (db *gameDb) UpdateGame(game Game) error {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `UPDATE games SET name = $1, description = $2, genre_id = $3, publisher_id = $4, platform_id = $5 WHERE game_id = $6`
	_, err := db.db.Exec(ctx, queryRow, game.Name, game.Description, game.GenreID, game.PublisherID, game.PlatformID, game.ID)
	return err
}

func (db *gameDb) DeleteGame(id string) error {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `DELETE FROM games WHERE game_id = $1`
	_, err := db.db.Exec(ctx, queryRow, id)
	return err
}


func NewGameDb(db *pgxpool.Pool) (IGame, error) {
	
	game := &gameDb{db: db}

	return game, nil
}