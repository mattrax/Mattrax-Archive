package checkin

import (
	"io"
	"log"
	"net/http"
)

func Handler(s interface{}) http.HandlerFunc {
	//Defines Structs/Static Stuff Here
	msg := "Checkin handler"

	return func(w http.ResponseWriter, r *http.Request) {
		var cmd structs.CheckinCommand
		if err := plist.NewXMLDecoder(r.Body).Decode(&cmd); err != nil {
			return 403, err
		}

		log.Println("Checkin")
		io.WriteString(w, msg)
	}
}
