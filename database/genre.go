package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type genreDb struct {
	db *pgxpool.Pool
}

type Genre struct {
	ID        int       `db:"genre_id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

type IGenre interface {
	GetGenres() ([]Genre, error)
	GetGenre(id int) (Genre, error)
	CreateGenre(genre Genre) (string, error)
	UpdateGenre(genre Genre) error
	DeleteGenre(id string) error
}

func (db *genreDb) GetGenres() ([]Genre, error) {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `SELECT genre_id, name, created_at FROM genres`
	rows, err := db.db.Query(ctx, queryRow)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genres []Genre
	for rows.Next() {
		var genre Genre
		if err := rows.Scan(&genre.ID, &genre.Name, &genre.CreatedAt); err != nil {
			return nil, err
		}
		genres = append(genres, genre)
	}
	return genres, nil
}

func (db *genreDb) GetGenre(id int) (Genre, error) {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `SELECT genre_id, name, created_at FROM genres WHERE genre_id = $1`
	row := db.db.QueryRow(ctx, queryRow, id)
	var genre Genre
	if err := row.Scan(&genre.ID, &genre.Name, &genre.CreatedAt); err != nil {
		return genre, err
	}
	return genre, nil
}

func (db *genreDb) CreateGenre(genre Genre) (string, error) {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `INSERT INTO genres (name) VALUES ($1)`
	_, err := db.db.Exec(ctx, queryRow, genre.Name)
	return "game created", err
}

func (db *genreDb) UpdateGenre(genre Genre) error {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `UPDATE genres SET name = $1 WHERE genre_id = $2`
	_, err := db.db.Exec(ctx, queryRow, genre.Name, genre.ID)
	return err
}

func (db *genreDb) DeleteGenre(id string) error {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `DELETE FROM genres WHERE genre_id = $1`
	_, err := db.db.Exec(ctx, queryRow, id)
	return err
}

func NewGenreDb(db *pgxpool.Pool) (IGenre, error) {
	
	genre := &genreDb{db: db}

	return genre, nil
}
