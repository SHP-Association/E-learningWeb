package log

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/context"
)

// Set sets a logger in the context.
func Set(ctx echo.Context, logger *slog.Logger) {
	ctx.Set(context.LoggerKey, logger)
}

// Ctx returns the logger stored in context, or provides the default logger if one is not present.
func Ctx(ctx echo.Context) *slog.Logger {
	if ctx == nil {
		return Default()
	}

	logger := Default()
	func() {
		defer func() {
			_ = recover()
		}()
		if l, ok := ctx.Get(context.LoggerKey).(*slog.Logger); ok {
			logger = l
		}
	}()

	return logger
}

// Default returns the default logger.
func Default() *slog.Logger {
	return slog.Default()
}
