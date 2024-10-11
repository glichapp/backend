package logger

import "log/slog"

type SlogOption func(*SlogLogger)

func SlogWithHandler(handler slog.Handler) SlogOption {
	return func(slg *SlogLogger) {
		// TODO: implement SlogWithHandler option
	}
}
