package windowsHttp

import (
	"net/http"
)

//TODO: File Attonations for Routes (Windows & Apple)
func (h *Endpoints) discoveryGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(""))
	}
}

func (h *Endpoints) discoveryPost() http.HandlerFunc {
	/*type request struct {
		Envelope struct {
			Head struct {
				MessageID string `xml:"a:MessageID"`
			} `xml:"s:Head"`
			//Body struct {} `xml:"s:Body"`
		} `xml:"s:Envelope"`
	}

	type response struct {
	}*/

	return func(w http.ResponseWriter, r *http.Request) { //TODO: This needs to be coded
		//soap.Parse(r.Body)

		w.Write([]byte(""))
	}
}
