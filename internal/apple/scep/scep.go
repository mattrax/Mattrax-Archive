package scep

import (
	"context"
	"io/ioutil"
	"net/http"

	scepserver "github.com/mattrax/Mattrax/pkg/scep"
	"github.com/mattrax/scep/depot/file" // TODO: Get Rid Of This One
)

var svc scepserver.Service

func init() {
	depot, err := file.NewFileDepot("depot") //TODO: Get From Config
	if err != nil {
		panic(err)
	}

	svcOptions := []scepserver.ServiceOption{ //TODO: ALlow All Of This To Be Configured
		scepserver.ChallengePassword("secret"), // The SCEP ChallengePassword
		//scepserver.WithCSRVerifier(csrVerifier), //TODO: Make This Work
		scepserver.CAKeyPassword([]byte("password")), // The CA Key's Password
		scepserver.ClientValidity(365),
		scepserver.AllowRenewal(0),
		//scepserver.WithLogger(logger), //TODO: Try Normal Logger
	}

	svc, err = scepserver.NewService(depot, svcOptions...)
	if err != nil {
		panic(err)
	}
}

// The HTTP Get Web Handler
func GetHandler() func(w http.ResponseWriter, r *http.Request) error {
	loadScepCA()

	return func(w http.ResponseWriter, r *http.Request) error {
		operation := r.URL.Query().Get("operation")

		switch operation {
		case "GetCACaps":
			res, err := svc.GetCACaps(context.Background())
			if err != nil {
				return err
			}
			w.Write(res)
		case "GetCACert":
			res, _, err := svc.GetCACert(context.Background()) //TODO: Handle The Middle Value
			if err != nil {
				return err
			}
			w.Header().Set("Content-Type", "application/x-x509-ca-cert")
			w.Write(res)
		case "GetNextCACert":
			res, err := svc.GetNextCACert(context.Background())
			if err != nil {
				return err
			}
			w.Write(res)
		default:
			http.Error(w, "Invalid Operation", 500)
		}

		return nil
	}
}

// The HTTP Post Web Handler
func PostHandler() func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		operation := r.URL.Query().Get("operation")

		if operation == "PKIOperation" {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				return err
			}
			res, err := svc.PKIOperation(context.Background(), body)
			if err != nil {
				return err
			}
			w.Write(res)
		} else {
			http.Error(w, "Invalid Operation", 500)
		}

		return nil
	}
}

// This Function Loads The SCEP Certificate From Disk And If It Is Not Found It Generates One.
func loadScepCA() {
	//Check If It Exists If So
	//   Load The Cert
	// Else
	//   Generate A Cert initScepCA()
}

// This Function Generates A New CA/Key For SCEP Signing
func initScepCA() {
	//Create A New CA
}
