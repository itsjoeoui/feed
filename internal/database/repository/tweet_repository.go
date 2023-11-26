package repository

import (
	"context"
	"database/sql"
	"feed/internal/domain/entity"
	r "feed/internal/ports/repository"
)

type tweetRepository struct {
	db *sql.DB
}

// Create implements ports.TweetRepository.
func (*tweetRepository) Create(ctx context.Context, tweet *entity.Tweet) (int error) {
	panic("unimplemented")
}

// Delete implements ports.TweetRepository.
func (*tweetRepository) Delete(ctx context.Context, id int) error {
	panic("unimplemented")
}

// Get implements ports.TweetRepository.
func (*tweetRepository) Get(ctx context.Context, id int) (*entity.Tweet, error) {
	panic("unimplemented")
}

// List implements ports.TweetRepository.
func (*tweetRepository) List(ctx context.Context) ([]*entity.Tweet, error) {
	panic("unimplemented")
}

// Search implements ports.TweetRepository.
func (*tweetRepository) Search(ctx context.Context, query string) ([]*entity.Tweet, error) {
	panic("unimplemented")
}

func NewTweetRepository(db *sql.DB) r.TweetRepository {
	return &tweetRepository{db: db}
}
