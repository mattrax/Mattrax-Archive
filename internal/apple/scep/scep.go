package scep

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mattrax/Mattrax/internal/server"
)

func Handler(server server.Server) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("Hello World!"))
	}
}
