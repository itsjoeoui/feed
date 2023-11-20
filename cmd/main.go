package main

import (
	"context"
	"feed/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	config, err := config.LoadAppConfig(".")
	if err != nil {
		// We are using zerolog
		// https://github.com/rs/zerolog
		log.Fatal().Err(err).Msg("unable to load configuration")
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	httplog.Configure(httplog.Options{
		Concise:         true,
		TimeFieldFormat: time.DateTime,
	})

	router := chi.NewRouter()

	router.Use(httplog.RequestLogger(log.Logger))
	router.Use(middleware.Timeout(60 * time.Second))
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)

	router.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	server := newServer(config.ServeAddress, router)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("unable to start server")
		}
	}()

	waitForShutdown(server)
}

func newServer(addr string, r *chi.Mux) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}

func waitForShutdown(server *http.Server) {
	// How does this function work?

	// We first deflare a Go channel named sig to receive a os.Signal
	sig := make(chan os.Signal, 1)
	// What Notify does is that it registers the signals to the channel
	// In this case, it registers Interrupt and SIGTERM
	// In this case, when one of these 2 gets triggered, it will send a signal to the channel
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig // this blocks the executation until the signal is received

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // making sure that we cancel the context

	// Very nice graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("server shutdown failed")
	}
}
