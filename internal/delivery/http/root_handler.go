package handler

import (
	repoErr "feed/internal/database/repository"
	uc "feed/internal/ports/usecase"
	"feed/internal/templates/pages"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type rootHandler struct {
	TweetUsecase uc.TweetUsecase
}

func NewRootHandler(r *chi.Mux, useCase uc.TweetUsecase) {
	handler := &rootHandler{
		TweetUsecase: useCase,
	}

	r.Route("/", func(r chi.Router) {
		r.Get("/", handler.ListTweets)
	})
}

func (h *rootHandler) ListTweets(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ts, err := h.TweetUsecase.ListTweets(ctx)
	if err != nil {
		log.Error().Msg(err.Error())
		select {
		case <-ctx.Done():
			http.Error(w, timeout, http.StatusGatewayTimeout)

		default:
			if err == repoErr.ErrListTweets {
				http.Error(w, listTweets, http.StatusNotFound)
			} else {
				http.Error(w, getTweet, http.StatusInternalServerError)
			}
		}
		return
	}

	templ.Handler(pages.HomePage(ts)).ServeHTTP(w, r)
}
