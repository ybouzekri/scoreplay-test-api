package main

import (
	"flag"
	"log/slog"
	"os"
	"scoreplay/internal/drivers/rest"
)

type appLogger struct {
	level string
}

func (l *appLogger) Level() slog.Level {
	switch l.level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func main() {
	var serverAddr string
	var logLevel string

	flag.StringVar(&serverAddr, "addr", "0.0.0.0:8888", "the server addr")
	flag.StringVar(&logLevel, "log-level", "info", "the logger verbosity (accepts debug, info, warning and error)")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: &appLogger{level: logLevel},
	}))

	logger.Info("logger initialized", "level", logLevel)

	router := rest.NewRouter(logger)
	server := rest.NewServer(serverAddr, router, logger)

	if err := server.ListenAndServe(); err != nil {
		logger.Error("server exited", "msg", err.Error())
		os.Exit(1)
	}
}
