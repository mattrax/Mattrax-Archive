package server

import (
	"io"
	"net/http"
)

func Handler() http.HandlerFunc {
	//Defines Structs/Static Stuff Here
	msg := "Hello World Server"
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, msg)
	}
}
