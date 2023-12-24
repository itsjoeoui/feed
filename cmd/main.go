package main

import (
	"context"
	"database/sql"
	"feed/config"
	"feed/internal/assets"
	r "feed/internal/database/repository"
	handler "feed/internal/delivery/http"
	u "feed/internal/domain/usecase"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
	"github.com/go-chi/jwtauth/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	config, err := config.LoadAppConfig()
	if err != nil {
		// We are using zerolog
		// https://github.com/rs/zerolog
		log.Fatal().Err(err).Msg("unable to load configuration")
	}

	// Setup for JWT
	tokenAuth := jwtauth.New("HS256", []byte(config.JWTSecret), nil)

	// Setup Google OAuth
	googleOAuthConfig := &oauth2.Config{
		RedirectURL:  config.RedirectURL,
		ClientID:     config.GoogleClientID,
		ClientSecret: config.GoogleClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	// Setup Postgres
	db, err := setupDB(config.DBDriver, config.DBURL)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to connect to the database")
	}
	defer db.Close()

	// Configure out logger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	httplog.Configure(httplog.Options{
		Concise:         true,
		TimeFieldFormat: time.DateTime,
	})

	// Setup the chi router
	router := chi.NewRouter()

	// Include middlewares that we want to enable for all routes.
	router.Use(httplog.RequestLogger(log.Logger))
	router.Use(middleware.Timeout(60 * time.Second))
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)

	// Mount everything in ./internal/assets/dist
	// This includes the static CSS file for example.
	assets.Mount(router)

	tweetRepo := r.NewTweetRepository(db)
	tweetUC := u.NewTweetUseCase(tweetRepo)

	handler.NewTweetHandler(router, tweetUC, tokenAuth)
	handler.NewRootHandler(router, tweetUC, tokenAuth)
	handler.NewAuthHandler(router, googleOAuthConfig, tokenAuth)

	// Start our server
	server := newServer(config.ServeAddress+":"+config.ServePort, router)

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

// setupDB initiates the database connection
func setupDB(driver, url string) (*sql.DB, error) {
	db, err := sql.Open(driver, url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
