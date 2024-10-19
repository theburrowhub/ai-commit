// Package logger provides functionality to set up the logger for the application.
package logger

import (
	"fmt"
	"log/slog"
	"os"
	"runtime/debug"
	"sync"
)

var ErrorLogger *slog.Logger

// SetupLogger sets up the logger for the application
func SetupLogger(logLevel string) {
	var level = setLevel(logLevel)
	var addSource = false
	var attributes sync.Map

	// Add source to debug logs and set custom attributes
	if level == slog.LevelDebug {
		addSource = true

		buildInfo, _ := debug.ReadBuildInfo()
		if buildInfo == nil {
			buildInfo = &debug.BuildInfo{}
		}
		attributes.Store("pid", fmt.Sprintf("%d", os.Getpid()))
		attributes.Store("go_version", buildInfo.GoVersion)
	}

	// Handlers for stdout and stderr
	stdoutHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     level,
		AddSource: addSource,
	})
	// Create a new logger for stdout
	logger := slog.New(stdoutHandler)
	// Set the logger as the default logger
	slog.SetDefault(logger)

	// Custom attributes for the error logger
	stderrHandler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level:     slog.LevelError,
		AddSource: addSource,
	})
	// Create a new logger for errors to stderr
	ErrorLogger = slog.New(stderrHandler)
}

func setLevel(logLevel string) slog.Level {
	var level slog.Level

	switch logLevel {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	return level
}
