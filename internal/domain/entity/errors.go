package entity

import "errors"

var ErrInvalidTweet = errors.New("content and user ID cannot be empty")
