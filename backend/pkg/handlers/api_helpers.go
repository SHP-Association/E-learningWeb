package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func parseIntParam(ctx echo.Context, name string) (int, error) {
	id, err := strconv.Atoi(ctx.Param(name))
	if err != nil || id <= 0 {
		return 0, echo.NewHTTPError(http.StatusBadRequest, "invalid "+name)
	}
	return id, nil
}

func jsonInternalError(ctx echo.Context, message string) error {
	return ctx.JSON(http.StatusInternalServerError, echo.Map{
		"error": message,
	})
}
