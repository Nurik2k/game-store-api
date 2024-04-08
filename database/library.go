package database

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type libraryDb struct {
	db *pgxpool.Pool
}

type Library struct {
	ID        string `db:"library_id"`
	GameID    int `db:"game_id"`
	UserID    int `db:"user_id"`
	CreatedAt int `db:"created_at"`
}

type ILibrary interface {
	GetLibrariesByUser(userID string) ([]Library, error)
	CreateLibrary(library Library) error
	AddGameToLibraryFromUser(library Library) error
	DeleteLibrary(id string) error
}

func (db *libraryDb) GetLibrariesByUser(userID string) ([]Library, error) {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `SELECT library_id, game_id, user_id, created_at FROM libraries WHERE user_id = $1`
	rows, err := db.db.Query(ctx, queryRow, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var libraries []Library
	for rows.Next() {
		var library Library
		if err := rows.Scan(&library.ID, &library.GameID, &library.UserID, &library.CreatedAt); err != nil {
			return nil, err
		}
		libraries = append(libraries, library)
	}
	return libraries, nil
}

func (db *libraryDb) CreateLibrary(library Library) error {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `INSERT INTO libraries (game_id, user_id) VALUES ($1, $2)`
	_, err := db.db.Exec(ctx, queryRow, library.GameID, library.UserID)
	return err
}

func (db *libraryDb) AddGameToLibraryFromUser(library Library) error {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `INSERT INTO libraries (game_id, user_id) VALUES ($1, $2)`
	_, err := db.db.Exec(ctx, queryRow, library.GameID, library.UserID)
	return err
}

func (db *libraryDb) DeleteLibrary(id string) error {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `DELETE FROM libraries WHERE library_id = $1`
	_, err := db.db.Exec(ctx, queryRow, id)
	return err
}

func NewLibraryDb(db *pgxpool.Pool) (ILibrary, error) {
	
	library := &libraryDb{db: db}

	return library, nil
}