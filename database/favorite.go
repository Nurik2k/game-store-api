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
	ID      string       `db:"favorite_id"`
	UserID  int       `db:"user_id"`
	GameID  int       `db:"game_id"`
	AddedAt time.Time `db:"added_at"`
}

type IFavoriteGame interface {
	GetFavoriteGamesByUser(userID string) ([]FavoriteGame, error)
	AddFavoriteGameToUser(favorite FavoriteGame) error
	DeleteFavoriteGameFromUser(id string) error
}

func (db *favoriteDb) GetFavoriteGamesByUser(userID string) ([]FavoriteGame, error) {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `SELECT fav_id, user_id, game_id, added_at FROM favorites WHERE user_id = $1`
	rows, err := db.db.Query(ctx, queryRow, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var favorites []FavoriteGame
	for rows.Next() {
		var favorite FavoriteGame
		if err := rows.Scan(&favorite.ID, &favorite.UserID, &favorite.GameID, &favorite.AddedAt); err != nil {
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

func (db *favoriteDb) DeleteFavoriteGameFromUser(id string) error {
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