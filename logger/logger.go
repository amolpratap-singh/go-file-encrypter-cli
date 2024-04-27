package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var log zerolog.Logger

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log = zerolog.New(os.Stdout).With().Timestamp().Logger()
}

// Info logs an informational message
func Info(msg string) {
    log.Info().
        Str("severity", "info").
        Str("timestamp", time.Now().Format(time.RFC3339)).
        Str("msg", msg).
        Msg("")
}

// Warn logs a warning message
func Warn(msg string) {
    log.Warn().
        Str("severity", "warning").
        Str("timestamp", time.Now().Format(time.RFC3339)).
        Str("msg", msg).
        Msg("")
}

// Error logs an error message
func Error(msg string) {
    log.Error().
        Str("severity", "error").
        Str("timestamp", time.Now().Format(time.RFC3339)).
        Str("msg", msg).
        Msg("")
}

// Panic logs a panic message and recovers from the panic
func Panic(msg string) {
    if r := recover(); r != nil {
        log.Error().
            Str("severity", "panic").
            Str("timestamp", time.Now().Format(time.RFC3339)).
            Str("msg", msg).
            Msg("panic recovered")
    }
}