package handler

const (
	timeout            = "request timed out"
	invalidRequestBody = "the request body is invalid or malformed"
)

const (
	getTweet       = "failed to retrieve the tweet"
	deleteTweet    = "failed to delete the tweet"
	createTweet    = "failed to create the tweet"
	tweetNotFound  = "the requested tweet was not found"
	invalidTweetID = "invalid tweet ID provided, it should be a positive integer"
)
