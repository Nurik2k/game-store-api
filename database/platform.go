package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type platformDb struct {
	db *pgxpool.Pool
}

type Platform struct {
	ID        int       `db:"platform_id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

type IPlatform interface {
	GetPlatforms() ([]Platform, error)
	GetPlatform(id int) (Platform, error)
	CreatePlatform(platform Platform) error
	UpdatePlatform(platform Platform) error
	DeletePlatform(id int) error
}

func (db *platformDb) GetPlatforms() ([]Platform, error) {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `SELECT platform_id, name, created_at FROM platforms`
	rows, err := db.db.Query(ctx, queryRow)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var platforms []Platform
	for rows.Next() {
		var platform Platform
		if err := rows.Scan(&platform.ID, &platform.Name, &platform.CreatedAt); err != nil {
			return nil, err
		}
		platforms = append(platforms, platform)
	}
	return platforms, nil
}

func (db *platformDb) GetPlatform(id int) (Platform, error) {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `SELECT platform_id, name, created_at FROM platforms WHERE platform_id = $1`
	row := db.db.QueryRow(ctx, queryRow, id)
	var platform Platform
	if err := row.Scan(&platform.ID, &platform.Name, &platform.CreatedAt); err != nil {
		return platform, err
	}
	return platform, nil
}

func (db *platformDb) CreatePlatform(platform Platform) error {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `INSERT INTO platforms (name) VALUES ($1)`
	_, err := db.db.Exec(ctx, queryRow, platform.Name)
	return err
}

func (db *platformDb) UpdatePlatform(platform Platform) error {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `UPDATE platforms SET name = $1 WHERE platform_id = $2`
	_, err := db.db.Exec(ctx, queryRow, platform.Name, platform.ID)
	return err
}

func (db *platformDb) DeletePlatform(id int) error {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `DELETE FROM platforms WHERE platform_id = $1`
	_, err := db.db.Exec(ctx, queryRow, id)
	return err
}

func NewPlatformDb(db *pgxpool.Pool) (IPlatform, error) {
	
	platform := &platformDb{db: db}
	
	return platform, nil
}
