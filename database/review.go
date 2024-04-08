package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type reviewDb struct {
	db *pgxpool.Pool
}

type Review struct {
	ID      string       `db:"review_id"`
	GameID  int       `db:"game_id"`
	UserID  int       `db:"user_id"`
	Rating  int       `db:"rating"`
	Review  string    `db:"review"`
	AddedAt time.Time `db:"added_at"`
}

type IReview interface {
	GetReviews() ([]Review, error)
	GetReview(id string) (Review, error)
	CreateReview(review Review) error
	UpdateReview(review Review) error
	DeleteReview(id string) error
}

func (db *reviewDb) GetReviews() ([]Review, error) {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `SELECT review_id, game_id, user_id, rating, review, added_at FROM reviews`
	rows, err := db.db.Query(ctx, queryRow)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []Review
	for rows.Next() {
		var review Review
		if err := rows.Scan(&review.ID, &review.GameID, &review.UserID, &review.Rating, &review.Review, &review.AddedAt); err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}

func (db *reviewDb) GetReview(id string) (Review, error) {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `SELECT review_id, game_id, user_id, rating, review, added_at FROM reviews WHERE review_id = $1`
	row := db.db.QueryRow(ctx, queryRow, id)
	var review Review
	if err := row.Scan(&review.ID, &review.GameID, &review.UserID, &review.Rating, &review.Review, &review.AddedAt); err != nil {
		return review, err
	}
	return review, nil
}

func (db *reviewDb) CreateReview(review Review) error {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `INSERT INTO reviews (game_id, user_id, rating, review) VALUES ($1, $2, $3, $4)`
	_, err := db.db.Exec(ctx, queryRow, review.GameID, review.UserID, review.Rating, review.Review)
	return err
}

func (db *reviewDb) UpdateReview(review Review) error {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `UPDATE reviews SET game_id = $1, user_id = $2, rating = $3, review = $4 WHERE review_id = $5`
	_, err := db.db.Exec(ctx, queryRow, review.GameID, review.UserID, review.Rating, review.Review, review.ID)
	return err
}

func (db *reviewDb) DeleteReview(id string) error {
	ctx := context.Background()
	defer ctx.Done()

	queryRow := `DELETE FROM reviews WHERE review_id = $1`
	_, err := db.db.Exec(ctx, queryRow, id)
	return err
}

func NewReviewDb(db *pgxpool.Pool) (IReview, error) {
	
	review := &reviewDb{db: db}

	return review, nil
}