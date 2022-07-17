package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Test1(ctx echo.Context) error {

	return ctx.JSON(http.StatusOK, "111")
}
