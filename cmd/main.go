package main

import (
	"net/http"
	"os"

	"github.com/sgkul2000/go-mailer/mail"
	"github.com/sgkul2000/go-mailer/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Debug = true

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${uri} ${status} ${latency_human}\n",
	}))

	e.HTTPErrorHandler = middlewares.EchoErrorHandler

	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {

		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/mail", mail.Mailer)
	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
