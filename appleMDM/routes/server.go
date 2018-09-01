package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

func ServerHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Server Route")
}
