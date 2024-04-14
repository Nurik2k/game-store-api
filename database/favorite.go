package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type favoriteDb struct {
	db *pgxpool.Pool
}

type FavoriteGame struct {
	ID      int       `db:"favorite_id"`
	UserID  int       `db:"user_id"`
	GameID  int       `db:"game_id"`
	AddedAt time.Time `db:"added_at"`
}

type IFavoriteGame interface {
	GetFavoriteGamesByUser(userID int) ([]Game, error)
	AddFavoriteGameToUser(favorite FavoriteGame) error
	DeleteFavoriteGameFromUser(id int) error
}

func (db *favoriteDb) GetFavoriteGamesByUser(userID int) ([]Game, error) {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `select g.name, g.description, ge.name as genre, p.name as publisher, pl.name as platform from games g
	join favorites_games f on f.game_id = g.game_id
	join genres ge on ge.genre_id = g.genre_id
	join publisher p on p.publisher_id = g.publisher_id
	join platforms pl on pl.platform_id = g.platform_id
	where f.user_id = $1`
	rows, err := db.db.Query(ctx, queryRow, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var favorites []Game
	for rows.Next() {
		var favorite Game
		if err := rows.Scan(&favorite.Name, &favorite.Description, &favorite.GenreName, &favorite.PublisherName, &favorite.PlatformName); err != nil {
			return nil, err
		}
		favorites = append(favorites, favorite)
	}
	return favorites, nil
}

func (db *favoriteDb) AddFavoriteGameToUser(favorite FavoriteGame) error {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `INSERT INTO favorites (user_id, game_id) VALUES ($1, $2)`
	_, err := db.db.Exec(ctx, queryRow, favorite.UserID, favorite.GameID)
	return err
}

func (db *favoriteDb) DeleteFavoriteGameFromUser(id int) error {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `DELETE FROM favorites WHERE fav_id = $1`
	_, err := db.db.Exec(ctx, queryRow, id)
	return err
}



func NewFavoriteDb(db *pgxpool.Pool) (IFavoriteGame, error) {
	
	favorite := &favoriteDb{db: db}
	
	return favorite, nil
}