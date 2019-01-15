package endpoints

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func serverHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Server Request")

		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "%q", dump)

		w.WriteHeader(http.StatusOK)
	}
}
