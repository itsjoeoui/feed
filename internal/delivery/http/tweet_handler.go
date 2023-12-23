package handler

import (
	"encoding/json"
	repoErr "feed/internal/database/repository"
	"feed/internal/domain/entity"
	uc "feed/internal/ports/usecase"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/rs/zerolog/log"
)

type tweetHandler struct {
	TweetUsecase uc.TweetUsecase
}

func NewTweetHandler(r *chi.Mux, useCase uc.TweetUsecase, tokenAuth *jwtauth.JWTAuth) {
	handler := &tweetHandler{
		TweetUsecase: useCase,
	}

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))

		r.Route("/v1/tweets", func(r chi.Router) {
			r.Get("/{id}", handler.GetTweet)
			r.Post("/", handler.CreateTweet)
			r.Delete("/{id}", handler.DeleteTweet)
		})
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
			if err == repoErr.ErrTweetNotFound {
				http.Error(w, tweetNotFound, http.StatusNotFound)
			} else {
				http.Error(w, getTweet, http.StatusInternalServerError)
			}
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

func (h *tweetHandler) CreateTweet(w http.ResponseWriter, r *http.Request) {
	var t entity.Tweet

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		log.Error().Msg(err.Error())
		http.Error(w, invalidRequestBody, http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	id, err := h.TweetUsecase.CreateTweet(ctx, &t)
	if err != nil {
		log.Error().Msg(err.Error())

		select {
		case <-ctx.Done():
			http.Error(w, timeout, http.StatusGatewayTimeout)
		default:
			http.Error(w, createTweet, http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(map[string]int{"id": id}); err != nil {
		log.Error().Msg(err.Error())
		http.Error(w, createTweet, http.StatusInternalServerError)
		return
	}
}

func (h *tweetHandler) DeleteTweet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Error().Msg(err.Error())
		http.Error(w, invalidTweetID, http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	err = h.TweetUsecase.DeleteTweet(ctx, id)
	if err != nil {
		log.Error().Msg(err.Error())

		select {
		case <-ctx.Done():
			http.Error(w, timeout, http.StatusGatewayTimeout)
		default:
			if err == repoErr.ErrTweetNotFound {
				http.Error(w, tweetNotFound, http.StatusNotFound)
			} else {
				http.Error(w, deleteTweet, http.StatusInternalServerError)
			}
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
