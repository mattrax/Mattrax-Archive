package main

import (
	"net/http"

	"./appleMDM/structs"         //TODO: Full Path
	devices "./internal/devices" //TODO: WTF Does This Need The Start To Import With Go-plus
	"./internal/pgsql"           //TODO: Full Path
	"github.com/labstack/echo"   //TODO: Full Path
)

var pgdb = pgsql.GetDB()

func apiRoutes(e *echo.Group) {
	e.GET("/device/:name", func(c echo.Context) error {
		device := devices.Computer{
			UUID:        c.Param("name"),
			DeviceState: &structs.MacOS_DeviceState{},
			DeviceInfo:  &structs.MacOS_DeviceInfo{},
		}

		if err := pgdb.Select(&device); err != nil && !pgsql.NotFound(err) {
			return err
		}

		return c.JSON(http.StatusOK, device)
	})
}
