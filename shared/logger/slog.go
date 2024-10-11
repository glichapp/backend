package logger

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"time"

	"github.com/kvizyx/glich/shared/config"
)

var (
	ErrInvalidLogLevel = errors.New("invalid logging level")
	ErrInvalidAppMode  = errors.New("invalid application mode")
)

var slogLevel = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
}

type SlogLogger struct {
	logger *slog.Logger
}

var _ StructuralLogger = &SlogLogger{}

type SlogOptions struct {
	AppMode config.AppMode
	Service string
	Level   string
}

func NewSlogLogger(opts SlogOptions, extraOpts ...SlogOption) (*SlogLogger, error) {
	var logger *slog.Logger

	level, ok := slogLevel[opts.Level]
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrInvalidLogLevel, opts.Level)
	}

	switch opts.AppMode {
	case config.AppModeProduction:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level:     level,
				AddSource: true,
			}),
		)
	case config.AppModeDevelopment:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				Level:     level,
				AddSource: true,
			}),
		)
	default:
		return nil, fmt.Errorf("%w: %s", ErrInvalidAppMode, opts.AppMode)
	}

	if len(opts.Service) != 0 {
		logger = logger.With(
			slog.String("service", opts.Service),
		)
	}

	// TODO: implement functional options for SlogLogger
	// for _, extraOpt := range extraOpts {
	// 	extraOpt(logger)
	// }

	return &SlogLogger{logger: logger}, nil
}

func (l *SlogLogger) Debug(input string, fields ...any) {
	l.handle(slog.LevelDebug, input, fields...)
}

func (l *SlogLogger) Info(input string, fields ...any) {
	l.handle(slog.LevelInfo, input, fields...)
}

func (l *SlogLogger) Warn(input string, fields ...any) {
	l.handle(slog.LevelWarn, input, fields...)
}

func (l *SlogLogger) Error(input string, fields ...any) {
	l.handle(slog.LevelError, input, fields...)
}

func (l *SlogLogger) Fatal(input string, fields ...any) {
	l.Error(input, fields...)
	os.Exit(1)
}

func (l *SlogLogger) With(args ...any) StructuralLogger {
	return &SlogLogger{logger: l.logger.With(args...)}
}

func (l *SlogLogger) handle(level slog.Level, input string, fields ...any) {
	handler := l.logger.Handler()

	if !handler.Enabled(context.Background(), level) {
		return
	}

	var pcs [1]uintptr
	runtime.Callers(3, pcs[:])

	record := slog.NewRecord(time.Now(), level, input, pcs[0])
	for _, field := range fields {
		record.Add(field)
	}

	_ = handler.Handle(context.Background(), record)
}
