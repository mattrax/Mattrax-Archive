package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

func CheckinHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Checkin Route")
}
