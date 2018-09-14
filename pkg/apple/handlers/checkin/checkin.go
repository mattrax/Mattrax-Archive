package checkin

import (
	"log"
	"net/http"

	"github.com/groob/plist"
)

func Handler(s interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cmd CheckinCommand
		if err := plist.NewXMLDecoder(r.Body).Decode(&cmd); err != nil {
			log.Println(err) //Return HTTP Error 403
		}

    switch e.Command.MessageType {
	   case "Authenticate":
       
     case "TokenUpdate":

     case "Checkout" //TODO: Check This

     default:
       //TODO: Return Error
   }

		log.Println(cmd)
		w.WriteHeader(200)
		w.Write([]byte(""))
	}
}
