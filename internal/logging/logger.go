package logging

import (
	"io"
	"os"
	"time"

	"github.com/crazy-max/swarm-cronjob/internal/model"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Configure configures logger
func Configure(fl *model.Flags, location *time.Location) {
	var err error
	var w io.Writer

	zerolog.TimestampFunc = func() time.Time {
		return time.Now().In(location)
	}

	if !fl.LogJSON {
		w = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC1123,
		}
	} else {
		w = os.Stdout
	}

	log.Logger = zerolog.New(w).With().Timestamp().Logger()

	logLevel, err := zerolog.ParseLevel(fl.LogLevel)
	if err != nil {
		log.Fatal().Err(err).Msgf("Unknown log level")
	} else {
		zerolog.SetGlobalLevel(logLevel)
	}
}
