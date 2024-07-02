package main

import (
	"Logging/fileLogger"

	"github.com/pkg/errors"
	// "github.com/rs/zerolog"
	// log "github.com/rs/zerolog/log"
	// "github.com/rs/zerolog/pkgerrors"
)

func main() {

	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	fileLogger := fileLogger.NewFileLogger()

	// log.Logger.Println("test")

	fileLogger.Println("test")

	fileLogger.Info().Str("Category", "Search").
		Int("DurationTime", 80).
		Msg("search log")

	err := f3()
	fileLogger.Error().Stack().Err(err).Msg("Manual Error func3()")
}

func f1() error {
	return errors.New("Manual Error")
}

func f2() error {
	err := f1()
	return err
}

func f3() error {
	err := f2()
	return err
}
