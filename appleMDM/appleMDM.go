package appleMDM

import (
	"log" //Maybe TEMP

	"./routes" //TODO: Make This Full Path
	"github.com/labstack/echo"
)

func Startup(e *echo.Group) {
	log.Println("Loaded The AppleMDM")

	// Routes
	e.GET("/enroll", routes.EnrollHandler)
	e.GET("/checkin", routes.CheckinHandler)
	e.GET("/server", routes.ServerHandler)
	//TODO: SCEP Handler

	//API

}

func Shutdown() {

}
