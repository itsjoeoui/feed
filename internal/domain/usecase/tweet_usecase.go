package usecase

import (
	"context"
	"feed/internal/domain/entity"
	r "feed/internal/ports/repository"
	u "feed/internal/ports/usecase"
)

type tweetUseCase struct {
	tweetRepo r.TweetRepository
}

func NewTweetUseCase(repository r.TweetRepository) u.TweetUsecase {
	return &tweetUseCase{
		tweetRepo: repository,
	}
}

// CreateTweet implements ports.TweetUsecase.
func (s *tweetUseCase) CreateTweet(ctx context.Context, tweet *entity.Tweet) (int, error) {
	panic("unimplemented")
}

// DeleteTweet implements ports.TweetUsecase.
func (s *tweetUseCase) DeleteTweet(ctx context.Context, id int) error {
	panic("unimplemented")
}

// GetTweet implements ports.TweetUsecase.
func (s *tweetUseCase) GetTweet(ctx context.Context, id int) (*entity.Tweet, error) {
	panic("unimplemented")
}

// ListTweets implements ports.TweetUsecase.
func (s *tweetUseCase) ListTweets(ctx context.Context) ([]*entity.Tweet, error) {
	panic("unimplemented")
}

// SearchTweets implements ports.TweetUsecase.
func (s tweetUseCase) SearchTweets(ctx context.Context, query string) ([]*entity.Tweet, error) {
	panic("unimplemented")
}
