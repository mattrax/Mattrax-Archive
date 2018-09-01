package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"time"

	"./appleMDM"       //TODO: Full Path
	"./internal/pgsql" //TODO: Full Path
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

/*
  Start Function
*/
func main() {
	// Create a new instance of Echo
	e := echo.New()

	e.HTTPErrorHandler = customHTTPErrorHandler
	e.GET("/error", func(c echo.Context) error {
		return errors.New("A custom error")
	})

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{ //TODO: Debugging/Development
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover()) //TODO: Debugging/Development

	// Routes
	interfaceRoutes(e)
	apiRoutes(e.Group("/api"))

	// Modules Startup
	appleMDM.Startup(e.Group("/apple"))

	// Start as a web server
	//TODO: Hide The Startup Message/Header
	//TODO: Disable Logging On Build (Go Generate Maybe)
	//TODO: Log To A File In Production
	//e.Logger.Fatal(e.Start(":8000")) // TODO: This Is For Faster Development Shutdowns
	//return

	//e.Logger.SetLevel(llog.INFO) //TODO: Load From A Config
	//e.Server.Addr = ":1323" //TODO: Load From A Config

	//Connect To The Database
	pgsql.Connect("postgres://oscar.beaumont:@localhost/mattrax_new?sslmode=disable")

	// Start the web server
	go func() {
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
	}
	// Modules Shutdown
	appleMDM.Shutdown()
	//Shutdown The Database
	pgsql.Disconnect()
}

/*
	Error Handler
*/
func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.Logger().Error(err)

	if code == 500 {
		c.String(code, "500 - The Server Encountered An Internal Error")
	} else if code == 404 {
		c.String(code, "404 - Nothing Was Found")
	}
	//TODO: Collect Alerts And Send Them Upstream To Me For Debugging/Fixing
}

//e.GET("/hello", func(c echo.Context) error { return c.String(http.StatusOK, "Hello, World!") })
//e.GET("/json", func(c echo.Context) error { return c.JSON(200, "{ 'hello': 'world' }") })

//TODO:
//	Database Constant Online Chekcing And Handle And Alert Admin On Failure and Error 500 Everything
//	Document How It Works On An Overaching Level On The Website
//	Make The MDM Modules Dynamiclly Loadable Using The Go "plugin" Interface
//  File Headers and Function/Struct Descriptions
//	HTTP Error Handling -> Genric DONT TELL THE CLIENT THE ERROR
//	Cisco SCEP Server Built In
//	Apple Update Server
//	Add Health Check (Database Connection Status, Internet Connectin Status) -> e.GET("/health", func(c echo.Context) error { return c.JSON(200, "{ 'hello': 'world' }") })
//  Compile Assets Into The Binary -> Load Some/All Into RAM For Preformance
//  https://echo.labstack.com/middleware/secure
//	Force Go Fmt'ed Code For All Pull Requests (Do TravisCI Test)
