package user

import (
	"Q2Bank/app"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Register(g *echo.Group, apps *app.ContainerApp) {
	h := &handler{
		apps: apps,
	}

	g.POST("/common-user", h.createCommonUser)
	g.POST("/seller-user", h.createSellerUser)
}

type handler struct {
	apps *app.ContainerApp
}

func (h *handler) createCommonUser(c echo.Context) error {
	var request CommonUser
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Falha ao recuperar os dados da requisição.")
	}

	if err := c.Validate(&request); err != nil {
		fmt.Println(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Faltam dados a serem informados na requisição.")
	}

	ctx := c.Request().Context()
	user, err := h.apps.User.Create(ctx, *request.ToUser())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *handler) createSellerUser(c echo.Context) error {
	var request SellerUser
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Falha ao recuperar os dados da requisição.")
	}

	if err := c.Validate(&request); err != nil {
		fmt.Println(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Faltam dados a serem informados na requisição.")
	}

	ctx := c.Request().Context()
	user, err := h.apps.User.Create(ctx, *request.ToUser())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}