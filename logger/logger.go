package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var Log zerolog.Logger

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	Log = zerolog.New(os.Stdout).With().Timestamp().Logger()
}
