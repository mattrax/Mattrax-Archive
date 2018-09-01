package appleMDM

import (
	"log" //Maybe TEMP

	"./routes" //TODO: Make This Full Path
	"github.com/labstack/echo"
)

func Startup(e *echo.Group) {
	log.Println("Loaded The AppleMDM")

	// Generate The Required Resources
	// TODO: Generate The MDM Enrollment Profile and etc From The Configs Details In The TEMP Dir

	// Routes
	e.GET("/enroll", routes.EnrollHandler)
	e.PUT("/checkin", routes.CheckinHandler)
	e.PUT("/server", routes.ServerHandler)
	//TODO: SCEP Handler

	//API

}

func Shutdown() {

}

/*
func init() {
	computer := MacOS{
		Computer: devices.Computer{
			Name: "Oscars Macbook Air",
		},
		//Name:      "My name",
		PushToken: "TestingPushToken",
	}

	log.Println(computer)
	log.Println(computer.Name)
	log.Println(computer.PushToken)
	log.Println(computer.Test())
}
*/
