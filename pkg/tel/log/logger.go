package log

import (
	"context"
	"log/slog"
)

// Loggable interface used to abstract the use of the log api.
// If necessary, use this interface for dependency injection, but direct use in the slog in a global way is also recommended.
type Loggable interface {
	Info(ctx context.Context, msg string, args ...any)
	Error(ctx context.Context, msg string, args ...any)
	Warn(ctx context.Context, msg string, args ...any)
	Fatal(ctx context.Context, msg string, args ...any)
}

type Logger struct {
	logger *slog.Logger
}

func New(logger *slog.Logger) *Logger {
	return &Logger{logger: logger}
}

func (l *Logger) Info(ctx context.Context, msg string, args ...any) {
	l.logger.InfoContext(ctx, msg, args...)
}

func (l *Logger) Warn(ctx context.Context, msg string, args ...any) {
	l.logger.WarnContext(ctx, msg, args...)
}

func (l *Logger) Error(ctx context.Context, msg string, args ...any) {
	l.logger.ErrorContext(ctx, msg, args...)
}

func (l *Logger) Fatal(ctx context.Context, msg string, args ...any) {
	l.logger.ErrorContext(ctx, "panic error: "+msg, args...)
	panic(msg)
}
