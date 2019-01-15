package endpoints

import (
	"net/http"

	"github.com/kataras/muxie"
	"github.com/mattrax/Mattrax/internal/certstore"
	"github.com/mattrax/Mattrax/internal/config"
	"github.com/mattrax/Mattrax/platform/apple/storage"
	"github.com/micromdm/scep/server"
)

// Endpoints contains the dependencies and handlers for the HTTP endpoints
type Service struct { // TODO: maybe Rename To Service
	SCEPService scepserver.Service
	Storage     storage.Service

	// TODO: Clean up below - May Not Even be Needed When Enrollment profile Code is removed
	CertStore          *certstore.CertStore // TEMP: certificates.Store
	PublicURL          string               // Protocol + Domain + Port
	TenantName         string
	ProfileDescription string // Optional
	SCEPChallenge      string
	Topic              string
}

// MountEndpoints mounts the http endpoints for the service
func (svc *Service) MountEndpoints(mux *muxie.Mux) {
	// TODO: Somewhere check that the Apple signed payload exists -> Read other docs for more info
	// TODO: Check The MDM Signature in a Middleware

	mux.Handle("/apple/trust", muxie.Methods().
		HandleFunc(http.MethodGet, trustHandler(svc)))

	mux.Handle("/apple/enroll", muxie.Methods().
		HandleFunc(http.MethodGet, enrollHandler(svc)))

	mux.Handle("/apple/scep", muxie.Methods().
		HandleFunc(http.MethodGet, scepHandler(svc.SCEPService)).
		HandleFunc(http.MethodPost, scepHandler(svc.SCEPService)))

	mux.Handle("/apple/checkin", muxie.Methods().
		HandleFunc(http.MethodPut, svc.authenticate(svc.checkinHandler())))

	mux.Handle("/apple/server", muxie.Methods().
		HandleFunc(http.MethodPut, serverHandler()))
}

// New creates and returns a new Endpoints
func New(config config.Config, certStore *certstore.CertStore, storage storage.Service) Service { // TODO: Verify The SCEP Challenge Password
	// depot, err := file.NewFileDepot("./depot") // TODO: Replace This With An Internal CA
	// if err != nil {
	// 	panic(err) // TEMP
	// }

	svcOptions := []scepserver.ServiceOption{ //TODO: Allow Some Of This To Be Configured/Managed Dynamicluy
		scepserver.ChallengePassword("secret"),
		//scepserver.WithCSRVerifier(csrVerifier), //TODO: Make This Work
		scepserver.CAKeyPassword([]byte("secret")),
		scepserver.ClientValidity(365),
		scepserver.AllowRenewal(0),
	}

	server, err := scepserver.NewService(certStore, svcOptions...)
	if err != nil {
		panic(err) // TEMP
	}

	// TODO: Return Error From This Func

	return Service{
		SCEPService:   server,
		Storage:       storage,
		CertStore:     certStore,
		PublicURL:     config.PublicURL,
		TenantName:    config.TenantName,
		SCEPChallenge: "secret",                                                      // TODO: Deal With This
		Topic:         "com.apple.mgmt.XServer.232a74b5-7a81-4d6c-82fa-00f351e2c4f9", // TODO: Deal With This
	}
}
