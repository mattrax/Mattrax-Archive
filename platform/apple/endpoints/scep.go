package endpoints

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/micromdm/scep/server"
)

func scepHandler(server scepserver.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		operation := r.URL.Query().Get("operation")
		ctx := context.Background()
		switch operation {
		case "GetCACaps":
			res, err := server.GetCACaps(ctx)
			if err != nil {
				panic(err)
			}
			w.Write(res)
		case "GetCACert":
			res, _, err := server.GetCACert(ctx) //TODO: Handle The Middle Value
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/x-x509-ca-cert")
			w.Write(res)
		case "PKIOperation":
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}
			res, err := server.PKIOperation(ctx, body)
			if err != nil {
				panic(err)
			}
			w.Write(res)
		case "GetNextCACert":
			res, err := server.GetNextCACert(ctx)
			if err != nil {
				panic(err)
			}
			w.Write(res)
		default:
			http.Error(w, "Invalid Operation", 500) // TODO: Is This the correct thing to respond
		}
	}
}
