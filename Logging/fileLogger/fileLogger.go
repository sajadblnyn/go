package fileLogger

import (
	"log"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

func NewFileLogger() *zerolog.Logger {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	file, err := os.OpenFile("log.json", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Println(err)
	}

	logger := zerolog.New(file).
		With().
		Timestamp().
		Str("AppName", "test_zero_log").
		Logger()

	return &logger

}
