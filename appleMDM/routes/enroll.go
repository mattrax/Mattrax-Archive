package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

func EnrollHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Enroll Route")
}
