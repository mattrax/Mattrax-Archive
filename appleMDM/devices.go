package appleMDM

import (
	"log"

	main "../"
)

func init() {
	computer := main.Computer{
		name: "My name",
	}

	log.Println(computer)
	log.Println(computer.name)
	log.Println(computer.test())
}
