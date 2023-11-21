package handler

import (
	"encoding/json"
	uc "feed/internal/ports/usecase"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type tweetHandler struct {
	TweetUsecase uc.TweetUsecase
}

func NewTweetHandler(r *chi.Mux, useCase uc.TweetUsecase) {
	handler := &tweetHandler{
		TweetUsecase: useCase,
	}

	r.Route("/v1/tweets", func(r chi.Router) {
		r.Get("/{id}", handler.GetTweet)
	})
}

func (h *tweetHandler) GetTweet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Error().Msg(err.Error())
		http.Error(w, invalidTweetID, http.StatusBadRequest)
	}

	ctx := r.Context()
	t, err := h.TweetUsecase.GetTweet(ctx, id)
	if err != nil {
		log.Error().Msg(err.Error())
		select {
		case <-ctx.Done():
			http.Error(w, timeout, http.StatusGatewayTimeout)

		default:
			// TODO: handler other cases

		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(t); err != nil {

		log.Error().Msg(err.Error())
		http.Error(w, getTweet, http.StatusInternalServerError)
	}
}
