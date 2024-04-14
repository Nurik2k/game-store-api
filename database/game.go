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
	ID            int       `db:"game_id"`
	Name          string    `db:"name"`
	Description   string    `db:"description"`
	GenresID      int     `db:"genres_id"`
	PublisherID   int       `db:"publisher_id"`
	PlatformsID   int     `db:"platforms_id"`
	CreatedAt     time.Time `db:"created_at"`
	GenreName     string    `db:"genre"`
	PublisherName string    `db:"publisher"`
	PlatformName  string    `db:"platform"`
}

type IGame interface {
	GetGames() ([]Game, error)
	GetGame(id int) (Game, error)
	GetGamesByGenre(genre string) ([]Game, error)
	CreateGame(game Game) error
	UpdateGame(game Game) error
	DeleteGame(id int) error
}

func (db *gameDb) GetGames() ([]Game, error) {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `select g.game_id, g.name, g.description, g.genres_id, g.publisher_id, g.platforms_id, g.created_at, ge.name as genre, p.name as publisher, pl.name as platform from games g
	join favorites_games f on f.game_id = g.game_id
	join genres ge on ge.genre_id = g.genres_id
	join publisher p on p.publisher_id = g.publisher_id
	join platforms pl on pl.platform_id = g.platforms_id`
	rows, err := db.db.Query(ctx, queryRow)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []Game
	for rows.Next() {
		var game Game
		if err := rows.Scan(&game.ID, &game.Name, &game.Description, &game.GenresID,
			&game.PublisherID, &game.PlatformsID, &game.CreatedAt, &game.GenreName,
			&game.PublisherName, &game.PlatformName); err != nil {
			return nil, err
		}
		games = append(games, game)
	}
	return games, nil
}

func (db *gameDb) GetGamesByGenre(genre string) ([]Game, error) {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `select g.game_id, g.name, g.description, g.genres_id, g.publisher_id, g.platforms_id, g.created_at, ge.name as genre, p.name as publisher, pl.name as platform 
	from games g
	join favorites_games f on f.game_id = g.game_id
	join genres ge on ge.genre_id = g.genres_id
	join publisher p on p.publisher_id = g.publisher_id
	join platforms pl on pl.platform_id = g.platforms_id
	where ge.name ilike $1`
	rows, err := db.db.Query(ctx, queryRow, genre)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []Game
	for rows.Next() {
		var game Game
		if err := rows.Scan(&game.ID, &game.Name, &game.Description, &game.GenresID,
			&game.PublisherID, &game.PlatformsID, &game.CreatedAt, &game.GenreName,
			&game.PublisherName, &game.PlatformName); err != nil {
			return nil, err
		}
		games = append(games, game)
	}
	return games, nil
}

func (db *gameDb) GetGame(id int) (Game, error) {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `SELECT game_id, name, description, genre_id, publisher_id, platform_id, created_at FROM games WHERE game_id = $1;`
	row := db.db.QueryRow(ctx, queryRow, id)
	var game Game
	if err := row.Scan(&game.ID, &game.Name, &game.Description, &game.GenresID, &game.PublisherID, &game.PlatformsID, &game.CreatedAt); err != nil {
		return game, err
	}
	return game, nil
}

func (db *gameDb) CreateGame(game Game) error {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `INSERT INTO games (name, description, genre_id, publisher_id, platform_id) VALUES ($1, $2, $3, $4, $5)`
	_, err := db.db.Exec(ctx, queryRow, game.Name, game.Description, game.GenresID, game.PublisherID, game.PlatformsID)
	return err
}

func (db *gameDb) UpdateGame(game Game) error {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `UPDATE games SET name = $1, description = $2, genre_id = $3, publisher_id = $4, platform_id = $5 WHERE game_id = $6`
	_, err := db.db.Exec(ctx, queryRow, game.Name, game.Description, game.GenresID, game.PublisherID, game.PlatformsID, game.ID)
	return err
}

func (db *gameDb) DeleteGame(id int) error {
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
