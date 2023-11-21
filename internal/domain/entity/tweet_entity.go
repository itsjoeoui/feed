package entity

import "time"

type Tweet struct {
	CreatedAt time.Time `json:"created_at"`
	Content   string    `json:"content"`
	UserID    int       `json:"user_id"`
	ID        int       `json:"id"`
}

func NewTweet(content string, userID int) (*Tweet, error) {
	tweet := &Tweet{
		CreatedAt: time.Now(),
		Content:   content,
		UserID:    userID,
	}

	if err := tweet.Validate(); err != nil {
		return nil, err
	}

	return tweet, nil
}

func (tweet *Tweet) Validate() error {
	if tweet.Content == "" || tweet.UserID <= 0 {
		return ErrInvalidTweet
	}
	return nil
}
