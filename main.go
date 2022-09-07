package main

import (
	"Q2Bank/api"
	"Q2Bank/app"
	"Q2Bank/store"
	"Q2Bank/utils/validator"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	e := echo.New()
	e.Validator = validator.NewValidator()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := sqlx.MustConnect("mysql", "root:12345678@tcp(localhost:3306)/Q2Bank?parseTime=true")
	store := store.NewContainerStore(db)

	app := app.NewContainerApp(app.Options{
		Stores: store,
	})

	api.Register(api.Options{
		Group: e.Group("/v1"),
		Apps:  app,
	})

	e.Logger.Fatal(e.Start(":1323"))
}
