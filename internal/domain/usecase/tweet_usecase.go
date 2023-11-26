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
func (s *tweetUseCase) CreateTweet(ctx context.Context, t *entity.Tweet) (int, error) {
	tweet, err := entity.NewTweet(t.Content, 69)
	if err != nil {
		return 0, err
	}

	id, err := s.tweetRepo.Create(ctx, tweet)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// DeleteTweet implements ports.TweetUsecase.
func (s *tweetUseCase) DeleteTweet(ctx context.Context, id int) error {
	err := s.tweetRepo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

// GetTweet implements ports.TweetUsecase.
func (s *tweetUseCase) GetTweet(ctx context.Context, id int) (*entity.Tweet, error) {
	tweet, err := s.tweetRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return tweet, nil
}

// ListTweets implements ports.TweetUsecase.
func (s *tweetUseCase) ListTweets(ctx context.Context) ([]*entity.Tweet, error) {
	panic("unimplemented")
}

// SearchTweets implements ports.TweetUsecase.
func (s tweetUseCase) SearchTweets(ctx context.Context, query string) ([]*entity.Tweet, error) {
	panic("unimplemented")
}
