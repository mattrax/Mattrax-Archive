package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

/*
  Start Function
*/
func main() {
	// Create a new instance of Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{ //TODO: Debugging/Development
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover()) //TODO: Debugging/Development

	// Routes
	routes(e)

	// Start as a web server
	//TODO: Log To A File In Production
	e.Logger.Fatal(e.Start(":8000")) // TODO: This Is For Faster Development Shutdowns
	return

	//e.Logger.SetLevel(llog.INFO) //TODO: Load From A Config
	//e.Server.Addr = ":1323" //TODO: Load From A Config

	// Start the server
	/*go func() {
		if err := e.Start(":8000"); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}*/
}

/*
  Webserver Routes
*/
func routes(e *echo.Echo) {
	//e.GET("/", func(c echo.Context) error { return c.File("html/index.html") })
	e.File("/", "html/index.html")
	//e.File("/favicon.ico", "images/favicon.ico")
	//e.Static("/static", "assets")

	e.GET("/hello", func(c echo.Context) error { return c.String(http.StatusOK, "Hello, World!") })
	e.GET("/json", func(c echo.Context) error { return c.JSON(200, "{ 'hello': 'world' }") })
}

/*
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
*/

//TODO:
//  File Header
//  Compile Assets Into The Binary -> Load Some/All Into RAM For Preformance
