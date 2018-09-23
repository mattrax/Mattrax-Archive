package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	// DB Connection
	const connURL = `postgres://oscar.beaumont:@localhost/oscar.beaumont` //TODO: Get The Connection From This
	db, err := sqlx.Connect("postgres", "user=oscar.beaumont dbname=oscar.beaumont sslmode=disable")
	if err != nil {
		log.Panic("Error Connecting To The Database", err)
	}
	defer db.Close() //TODO: Check This Is Being Called

	//TODO: Check Check/Init

	// Webserver
	r := mux.NewRouter()

	routes(r, db)

	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}
	srv := &http.Server{
		Handler: r,
		Addr:    ":8001",
		// Good practice: enforce timeouts for servers you create!
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
		// TLS Security
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}

	log.Println("Mattrax Started Listening At ") //TODO: Finish This Message

	/*go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		signal.Stop(quit)
		fmt.Println("Shutting Down")
		ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
		srv.Shutdown(ctx)
	}()*/

	err = srv.ListenAndServeTLS("certs/bundle.crt", "certs/server.key") //TODO: Load Location From Config
	if err != nil {
		log.Panic("App Failed to start with the Error:", err.Error())
	}
}

/*
go func() {
  quit := make(chan os.Signal, 1)
  signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
  <-quit
  signal.Stop(quit)
  fmt.Println("Shutting Down")
  ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
  server.Shutdown(ctx)
}()

log.Fatal(server.ListenAndServe()) //TODO: Only Run This If Error
*/
