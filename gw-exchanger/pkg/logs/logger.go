package logger

import (
	"io"
	"strings"

	"github.com/rs/zerolog"
)

type Logger struct {
	*zerolog.Logger
}

func New(writer io.Writer, logLevel string) *Logger {
	parsedLevel, err := zerolog.ParseLevel(strings.ToLower(logLevel))
	if err != nil {
		parsedLevel = zerolog.InfoLevel
	}
	logger := zerolog.New(writer).Level(parsedLevel).With().Timestamp().Logger()
	if err != nil {
		logger.Info().Msg("Logging level is set to info")
	}
	logger.Info().Msgf("Logging level is set to %s", parsedLevel)
	return &Logger{&logger}
}

func (l *Logger) Debug() *zerolog.Event {
	return l.Logger.Debug()
}

func (l *Logger) Info() *zerolog.Event {
	return l.Logger.Info()
}

func (l *Logger) Warn() *zerolog.Event {
	return l.Logger.Warn()
}

func (l *Logger) Error() *zerolog.Event {
	return l.Logger.Error()
}

func (l *Logger) Fatal() *zerolog.Event {
	return l.Logger.Fatal()
}

func (l *Logger) Err(err error) *zerolog.Event {
	return l.Logger.Err(err)
}
