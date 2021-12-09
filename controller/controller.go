package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// IndexGet .
func IndexGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "GET /")
}

// AppsGet .
func AppsGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "GET /apps")
}

// AppsPost .
func AppsPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "POST /apps")
}
