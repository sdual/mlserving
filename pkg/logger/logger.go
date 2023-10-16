package logger

import (
	"github.com/rs/zerolog"
)

func SetUpLogger(logLevel string) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(selectLogLevel(logLevel))
}

func selectLogLevel(logLevel string) zerolog.Level {
	switch logLevel {
	case "info":
		return zerolog.InfoLevel
	case "debug":
		return zerolog.DebugLevel
	case "error":
		return zerolog.ErrorLevel
	default:
		// TODO: default をどうするか考える
		return zerolog.InfoLevel
	}
}
