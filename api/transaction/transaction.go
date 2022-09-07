package transaction

import (
	"Q2Bank/app"
	"Q2Bank/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Register handlers /transaction
func Register(g *echo.Group, apps *app.ContainerApp) {
	h := &handler{
		apps: apps,
	}

	g.POST("", h.create)
}

type handler struct {
	apps *app.ContainerApp
}

func (h *handler) create(c echo.Context) error {
	var request model.Transaction
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Falha ao recuperar os dados da requisição.")
	}

	if err := c.Validate(&request); err != nil {
		fmt.Println(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Faltam dados a serem informados na requisição.")
	}

	ctx := c.Request().Context()
	transaction, err := h.apps.Transaction.Create(ctx, request)
	if err != nil {
		fmt.Println("h.apps.Transaction.Create: ", err.Error())
		return err
	}

	return c.JSON(http.StatusCreated, transaction)
}
