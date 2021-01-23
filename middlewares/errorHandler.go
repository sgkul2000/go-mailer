package middlewares

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

// EchoErrorHandler is An error handler for echo frameword
func EchoErrorHandler(err error, c echo.Context) {
	fmt.Println(err)
	c.Echo().DefaultHTTPErrorHandler(err, c)
}
