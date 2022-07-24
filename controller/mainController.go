package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Test1(ctx echo.Context) error {
	url := ctx.Path()
	return ctx.JSON(http.StatusOK, url+"111")
}
