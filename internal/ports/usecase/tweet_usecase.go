package ports

import (
	"context"
	"feed/internal/domain/entity"
)

type TweetUsecase interface {
	GetTweet(ctx context.Context, id int) (*entity.Tweet, error)
	ListTweets(ctx context.Context) ([]*entity.Tweet, error)
	SearchTweets(ctx context.Context, query string) ([]*entity.Tweet, error)
	CreateTweet(ctx context.Context, tweet *entity.Tweet) (int, error)
	DeleteTweet(ctx context.Context, id int) error
}
