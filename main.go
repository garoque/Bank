package main

import (
	"Q2Bank/api"
	"Q2Bank/app"
	_ "Q2Bank/docs"
	"Q2Bank/store"
	"Q2Bank/utils/validator"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/go-sql-driver/mysql"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Swagger Q2Bank
// @version 1.0
// @description This is a sample server Q2Bank.

// @BasePath /v1
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

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
