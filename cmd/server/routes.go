package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mattrax/scep/depot/file"
	scepserver "github.com/mattrax/scep/server"
)

func (server *Server) routes() {
	//Vue

	//MDM Endpoints
	server.router.PUT("/apple/checkin", CheckinHandler)
	server.router.PUT("/apple/server", ServerHandler)

	server.router.GET("/apple/scep", ScepHandler)
	server.router.POST("/apple/scep", ScepHandler)

	//TEMP
	server.router.GET("/", Index)
}

////// TEMP ///////

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Hello World!"))
}

func CheckinHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	/*body, _ := ioutil.ReadAll(r.Body)
	sig, err := base64.StdEncoding.DecodeString(r.Header.Get("Mdm-Signature")) //TODO: Handle Without The Sig Header And Using The Directory Cert
	if err != nil {
		panic(err)
	}
	p7, err := pkcs7.Parse([]byte(sig))
	if err != nil {
		panic(err)
	}
	p7.Content = body
	if err := p7.Verify(); err != nil {
		panic(err)
	}
	cert := p7.GetOnlySigner()
	if cert == nil {
		panic("Missing Signer")
	}*/

	log.Println("Checkin Req")
	//bodyBuffer, _ := ioutil.ReadAll(r.Body)
	//log.Println(string(bodyBuffer))
	w.Write([]byte(""))
}

func ServerHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//w.Write([]byte("Hello World!"))
}

func ScepHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	depot, err := file.NewFileDepot("depoty")
	if err != nil {
		panic(err)
	}

	svcOptions := []scepserver.ServiceOption{ //TODO: ALlow All Of This To Be Configured
		scepserver.ChallengePassword("secret"), // The SCEP ChallengePassword
		//scepserver.WithCSRVerifier(csrVerifier), //TODO: Make This Work
		scepserver.CAKeyPassword([]byte("")), // The CA Key's Password
		scepserver.ClientValidity(365),
		scepserver.AllowRenewal(0),
		//scepserver.WithLogger(logger),
	}

	svc, err := scepserver.NewService(depot, svcOptions...)
	if err != nil {
		panic(err)
	}

	log.Println("SCEP GET From ", r.Method, " To: ", r.URL.String())

	operation := r.URL.Query().Get("operation")

	switch operation {
	case "GetCACaps":
		res, err := svc.GetCACaps(context.Background())
		if err != nil {
			panic(err)
		}
		w.Write(res)
	case "GetCACert":
		res, _, err := svc.GetCACert(context.Background()) //TODO: Handle The Middle Value
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
		res, err := svc.PKIOperation(context.Background(), body)
		if err != nil {
			panic(err)
		}
		w.Write(res)
	case "GetNextCACert":
		res, err := svc.GetNextCACert(context.Background())
		if err != nil {
			panic(err)
		}
		w.Write(res)
	default:
		http.Error(w, "Invalid Operation", 500)
	}
}
