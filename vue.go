package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func interfaceRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error { return c.String(http.StatusOK, "Welcome To Mattrax!") })

	//e.GET("/", func(c echo.Context) error { return c.File("html/index.html") })
	//e.File("/", "html/index.html")
	//e.Static("/static", "assets")

	/*var fs http.FileSystem = http.Dir("assets")

	err := vfsgen.Generate(fs, vfsgen.Options{})
	if err != nil {
		log.Fatalln(err)
	}*/
}
