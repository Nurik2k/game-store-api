package database

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type publisherDb struct {
	db *pgxpool.Pool
}

type Publisher struct {
	ID        string    `db:"publisher_id"`
	Name      string `db:"name"`
	CreatedAt int    `db:"created_at"`
}

type IPublisher interface {
	GetPublishers() ([]Publisher, error)
	GetPublisher(id string) (Publisher, error)
	CreatePublisher(publisher Publisher) error
	UpdatePublisher(publisher Publisher) error
	DeletePublisher(id string) error
}

func (db *publisherDb) GetPublishers() ([]Publisher, error) {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `SELECT publisher_id, name, created_at FROM publishers`
	rows, err := db.db.Query(ctx, queryRow)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var publishers []Publisher
	for rows.Next() {
		var publisher Publisher
		if err := rows.Scan(&publisher.ID, &publisher.Name, &publisher.CreatedAt); err != nil {
			return nil, err
		}
		publishers = append(publishers, publisher)
	}
	return publishers, nil
}

func (db *publisherDb) GetPublisher(id string) (Publisher, error) {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `SELECT publisher_id, name, created_at FROM publishers WHERE publisher_id = $1`
	row := db.db.QueryRow(ctx, queryRow, id)
	var publisher Publisher
	if err := row.Scan(&publisher.ID, &publisher.Name, &publisher.CreatedAt); err != nil {
		return publisher, err
	}
	return publisher, nil
}

func (db *publisherDb) CreatePublisher(publisher Publisher) error {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `INSERT INTO publishers (name) VALUES ($1)`
	_, err := db.db.Exec(ctx, queryRow, publisher.Name)
	return err
}

func (db *publisherDb) UpdatePublisher(publisher Publisher) error {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `UPDATE publishers SET name = $1 WHERE publisher_id = $2`
	_, err := db.db.Exec(ctx, queryRow, publisher.Name, publisher.ID)
	return err
}

func (db *publisherDb) DeletePublisher(id string) error {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `DELETE FROM publishers WHERE publisher_id = $1`
	_, err := db.db.Exec(ctx, queryRow, id)
	return err
}

func NewPublisherDb(db *pgxpool.Pool) (IPublisher, error) {
	
	publisher := &publisherDb{db: db}
	
	return publisher, nil
}
