package user

import (
	"Q2Bank/app"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Register handlers /user
func Register(g *echo.Group, apps *app.ContainerApp) {
	h := &handler{
		apps: apps,
	}

	g.POST("/common-user", h.createCommonUser)
	g.POST("/seller-user", h.createSellerUser)
	g.POST("/cash-deposit", h.cashDeposit)
}

type handler struct {
	apps *app.ContainerApp
}

// user swagger document
// @Description Create common user
// @Param user body CommonUser true "add user"
// @Tags user
// @Accept json
// @Produce json
// @Success 201 {object} CommonUser
// @Failure 400
// @Failure 409
// @Failure 500
// @Router /v1/user/common-user [post]
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
		fmt.Println("h.apps.User.Create: ", err.Error())
		return err
	}

	return c.JSON(http.StatusCreated, user)
}

// user swagger document
// @Description Create seller user
// @Param user body SellerUser true "add user"
// @Tags user
// @Accept json
// @Produce json
// @Success 201 {object} SellerUser
// @Failure 400
// @Failure 409
// @Failure 500
// @Router /v1/user/seller-user [post]
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

// user swagger document
// @Description Deposit cash in account user
// @Param user body RequestCashDeposit true "add cash"
// @Tags user
// @Accept json
// @Produce json
// @Success 204
// @Failure 400
// @Failure 500
// @Router /v1/user/cash-deposit [post]
func (h *handler) cashDeposit(c echo.Context) error {
	var request RequestCashDeposit
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Falha ao recuperar os dados da requisição.")
	}

	if err := c.Validate(&request); err != nil {
		fmt.Println(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Faltam dados a serem informados na requisição.")
	}

	ctx := c.Request().Context()
	err := h.apps.User.IncreaseBalance(ctx, request.Value, request.UserID)
	if err != nil {
		fmt.Println("h.apps.User.IncreaseBalance: ", err.Error())
		return err
	}

	return c.JSON(http.StatusNoContent, nil)
}
