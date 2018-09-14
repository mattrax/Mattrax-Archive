package enroll

import (
	"io"
	"net/http"
)

func MDMServiceConfigHandler() http.HandlerFunc { // https://developer.apple.com/enterprise/documentation/MDM-Protocol-Reference.pdf Bottom of Page 215
	//Defines Structs/Static Stuff Here
	msg := "MDM Service Config Here"
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, msg)
	}
}
