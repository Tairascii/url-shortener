package internal

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Tairascii/url-shortener/db"
	"github.com/Tairascii/url-shortener/internal/handler"
	"github.com/Tairascii/url-shortener/internal/infra/config"
	"github.com/Tairascii/url-shortener/internal/repository"
	"github.com/Tairascii/url-shortener/internal/usecase"
)

type App struct {
	cfg *config.Config
}

func NewApp(cfg *config.Config) *App {
	return &App{
		cfg: cfg,
	}
}

func (a *App) Start() error {
	ctx := context.Background()
	postgresDB, err := db.ConnectPostgres(ctx, db.Settings{
		User:     a.cfg.DB.User,
		Password: a.cfg.DB.Password,
		Host:     a.cfg.DB.Host,
		Port:     a.cfg.DB.Port,
		DBName:   a.cfg.DB.DBName,
	})
	if err != nil {
		return err
	}

	repo := repository.New(postgresDB, a.cfg.DB.DCID, a.cfg.DB.ShardID)
	uc := usecase.New(repo)
	h := handler.New(uc)

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", a.cfg.Service.Port),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handler.AttachRoutes(h).ServeHTTP(w, r)
		}),
		ReadTimeout:  a.cfg.Service.ReadTimeout,
		WriteTimeout: a.cfg.Service.WriteTimeout,
		IdleTimeout:  a.cfg.Service.IdleTimeout,
	}

	go func() {
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("start server: %v\n", err)
		}
	}()

	log.Printf("listening on port: %s\n", a.cfg.Service.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-quit

	slog.Info("shutting down server")

	return srv.Shutdown(ctx)
}
