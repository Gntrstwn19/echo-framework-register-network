package route

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func InitRoute() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	return e
}
