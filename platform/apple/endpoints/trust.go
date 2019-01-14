package endpoints

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// TODO: This returns the Mattrax trust profile - For signing or maybe put it in the enrollment profile???

func trustHandler(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "%q", dump)
		fmt.Println("REQ")

	}
}
