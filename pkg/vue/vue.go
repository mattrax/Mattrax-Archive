package vue

import (
	"io"
	"net/http"
)

func IndexHandler() http.HandlerFunc {
	//Defines Structs/Static Stuff Here
	msg := "Hello World Vue"
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, msg)
	}
}
