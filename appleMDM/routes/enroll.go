package routes

import (
	"github.com/labstack/echo"
)

func EnrollHandler(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/x-apple-aspen-config")
	return c.File("enroll.mobileconfig")
}
