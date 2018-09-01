package routes

import (
	"github.com/labstack/echo"
)

func EnrollHandler(c echo.Context) error {
	//w.Header().Set("Content-Type", "application/x-apple-aspen-config");
	return c.File("../../data/enroll.moileconfig")
}
