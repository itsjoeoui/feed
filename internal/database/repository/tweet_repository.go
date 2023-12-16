package repository

import (
	"context"
	"database/sql"
	"feed/internal/domain/entity"
	r "feed/internal/ports/repository"
	"fmt"
)

type tweetRepository struct {
	db *sql.DB
}

// ListByUserID implements ports.TweetRepository.
func (*tweetRepository) ListByUserID(ctx context.Context) ([]*entity.Tweet, error) {
	panic("unimplemented")
}

// Create implements ports.TweetRepository.
func (r *tweetRepository) Create(ctx context.Context, tweet *entity.Tweet) (int, error) {
	stmt, err := r.db.PrepareContext(ctx, "INSERT INTO tweets (content, user_id) VALUES ($1, $2) RETURNING id")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", ErrPrepareStatement, err)
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, tweet.Content, tweet.UserID).Scan(&tweet.ID)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", ErrExecuteQuery, err)
	}

	return tweet.ID, nil
}

// Delete implements ports.TweetRepository.
func (r *tweetRepository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.PrepareContext(ctx, "DELETE FROM tweets WHERE id = $1")
	if err != nil {
		return fmt.Errorf("%s: %w", ErrPrepareStatement, err)
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return fmt.Errorf("%s: %w", ErrExecuteStatement, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", ErrRetrieveRows, err)
	}

	if rowsAffected == 0 {
		return ErrTweetNotFound
	}

	return nil
}

// Get implements ports.TweetRepository.
func (r *tweetRepository) Get(ctx context.Context, id int) (*entity.Tweet, error) {
	stmt, err := r.db.PrepareContext(ctx, "SELECT * FROM tweets WHERE id = $1")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrPrepareStatement, err)
	}
	defer stmt.Close()

	t := &entity.Tweet{}

	row := stmt.QueryRowContext(ctx, id)

	err = row.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrTweetNotFound
		} else {
			return nil, fmt.Errorf("%s, %w", ErrScanData, err)
		}
	}

	return t, nil
}

// List implements ports.TweetRepository.
func (r *tweetRepository) List(ctx context.Context) ([]*entity.Tweet, error) {
	stmt, err := r.db.PrepareContext(ctx, "SELECT * FROM tweets ORDER BY created_at DESC")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrPrepareStatement, err)
	}
	defer stmt.Close()

	ts := []*entity.Tweet{}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrExecuteQuery, err)
	}

	for rows.Next() {
		var t entity.Tweet
		err = rows.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", ErrScanData, err)
		}

		ts = append(ts, &t)
	}

	return ts, nil
}

// Search implements ports.TweetRepository.
func (*tweetRepository) Search(ctx context.Context, query string) ([]*entity.Tweet, error) {
	panic("unimplemented")
}

func NewTweetRepository(db *sql.DB) r.TweetRepository {
	return &tweetRepository{db: db}
}
