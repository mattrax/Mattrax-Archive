package main

// TODO: //go:generate yarn -s --prod --non-interactive --cwd ./../../interface/
// TODO: //go:generate yarn -s --prod --non-interactive --cwd ./../../interface/ build

import (
	"context"
	"crypto/tls"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	mattrax "github.com/mattrax/Mattrax/internal"
	"github.com/mattrax/Mattrax/internal/builtin"
	"github.com/mattrax/Mattrax/platform/windows"
)

var ( // TODO: Clean Up All The Descriptions
	flgHTTPSPort    = flag.Int("https.port", 8443, "the webservers https port")
	flgHTTPSCert    = flag.String("https.cert", "cert/key.pem", "the webservers https cert")
	flgHTTPSKey     = flag.String("https.key", "cert/key.key", "the webservers https key")
	flgDatabasePath = flag.String("database.path", "mattrax.db", "the Mattrax database")
)

type Mattrax struct {
	r           *mux.Router
	config      mattrax.Config
	DataStore   mattrax.DataStore
	AuthService mattrax.AuthService
	WindowsMDM  mattrax.MDM
	// TODO: Apple
}

func main() {
	flag.Parse()
	m := Mattrax{
		r: mux.NewRouter(),
		// TODO: config
		DataStore: builtin.NewDataStore(*flgDatabasePath),
	}
	m.AuthService = builtin.NewAuthService(m.DataStore.(builtin.UserDataStore)) // TOOD: This cast will probs need error handling
	m.WindowsMDM = windows.MDM(m.DataStore, m.AuthService)
	// TODO: Apple
	m.routes()

	// TODO: HTTPS Certs Locally or Lets Encrypt (Handle The Second Domain For Windows Maybe)
	srv := &http.Server{
		Addr:           ":" + strconv.Itoa(*flgHTTPSPort),
		Handler:        m.r,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		TLSConfig: &tls.Config{
			PreferServerCipherSuites: true,
			MinVersion:               tls.VersionTLS12,
			CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		},
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		signal.Notify(sigint, syscall.SIGTERM)

		<-sigint

		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal(err) // TODO: Correct Logger
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err) // TODO: Correct Logger
	}

	<-idleConnsClosed
}

// TODO: Automatic HTTPS
// TODO: Logging
// TODO: Gzip Response
// TODO: w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
// TOOD: ErrorLog http.Server
// TODO: Server header w.Header().Set("Server", "")
// TODO: Version Information
