package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/djwhocodes/ticket-reservation/configs"
	pgRepo "github.com/djwhocodes/ticket-reservation/internal/repository/postgres"
	redisRepo "github.com/djwhocodes/ticket-reservation/internal/repository/redis"
)

func main() {
	cfg := configs.Load()

	pgPool, err := pgRepo.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer pgPool.Close()

	redisClient, err := redisRepo.New(cfg.RedisAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer redisClient.Close()

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: mux,
	}

	go func() {
		slog.Info("server started", "port", cfg.AppPort)

		if err := server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	stop := make(chan os.Signal, 1)

	signal.Notify(
		stop,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	<-stop

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	slog.Info("shutting down")

	_ = server.Shutdown(ctx)
}

//        _____________NOTES_____________

// "context" = For: cancellation, timeouts, deadlines
// "log" = Simple logging
// "log/slog" = Structured logging (Go 1.21+).
// { slog.Info( "server started", "port", 8080,) => time = ... level=INFO msg="server started" port=8080}
// "net/http" = Go's HTTP server package. Provides: Server, Handler, Request, ResponseWriter
// "os" = Operating system functionality. os.Exit(1) or os.Signal
// "os/signal" = Receives OS signals. Example: Ctrl+C, kill
// "syscall" = Provides constants like: syscall.SIGINT, syscall.SIGTERM
// "time" = timeouts, sleep, durations

// if err != nil { log.Fatal(err) } = If DB connection fails: Print error, Terminate program, because app cannot function.
// defer pgPool.Close() => Don't close now. Close when main() exits.

// mux := http.NewServeMux() => Creates router. mux.HandleFunc("/movies", ...)
// server := &http.Server{}  => Creates HTTP server.
// Requests go to => Browser -> Server -> mux -> handler
