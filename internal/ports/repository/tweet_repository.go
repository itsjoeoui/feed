package ports

import (
	"context"
	"feed/internal/domain/entity"
)

type TweetRepository interface {
	Get(ctx context.Context, id int) (*entity.Tweet, error)
	List(ctx context.Context) ([]*entity.Tweet, error)
	Search(ctx context.Context, query string) ([]*entity.Tweet, error)
	Create(ctx context.Context, tweet *entity.Tweet) (int error)
	Delete(ctx context.Context, id int) error
}
