package api

import (
	"Q2Bank/app"

	"github.com/labstack/echo/v4"
)

type Options struct {
	Group *echo.Group
	Apps  *app.ContainerApp
}

func Register(opts Options) {
}
