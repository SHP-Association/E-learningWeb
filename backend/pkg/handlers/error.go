package handlers

import (
	"net/http"
	"strings"

	"github.com/SHP-Association/E-learningWeb/backend/pkg/context"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/log"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui/pages"
	"github.com/labstack/echo/v4"
)

type Error struct{}

func (e *Error) Page(err error, ctx echo.Context) {
	if ctx.Response().Committed || context.IsCanceledError(err) {
		return
	}

	// Determine the error status code.
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	// Log the error.
	logger := log.Ctx(ctx)
	switch {
	case code >= 500:
		logger.Error(err.Error())
	case code >= 400:
		logger.Warn(err.Error())
	}

	// Return JSON for API routes.
	if strings.HasPrefix(ctx.Path(), "/api") {
		msg := http.StatusText(code)
		if he, ok := err.(*echo.HTTPError); ok {
			switch m := he.Message.(type) {
			case string:
				msg = m
			case error:
				msg = m.Error()
			}
		}
		if e := ctx.JSON(code, echo.Map{"error": msg}); e != nil {
			log.Ctx(ctx).Error("failed to write api error response", "error", e)
		}
		return
	}

	// Set the status code.
	ctx.Response().WriteHeader(code)

	// Render the error page.
	if err = pages.Error(ctx, code); err != nil {
		log.Ctx(ctx).Error("failed to render error page",
			"error", err,
		)
	}
}
