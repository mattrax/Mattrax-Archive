package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"upper.io/db.v3/postgresql"
)

func main() {
	// DB Connection
	const connURL = `postgres://oscar.beaumont:@localhost/oscar.beaumont`
	settings, err := postgresql.ParseURL(connURL)
	db, err := postgresql.Open(settings) // Open a connection.
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Configuration
	var configuration struct {
		AdminDomain           string `db:"AdminDomain"`
		Domain                string `db:"Domain"`
		OrganisationName      string `db:"OrganisationName"`
		OrganisationShortName string `db:"OrganisationShortName"`
		OrganisationEmail     string `db:"OrganisationEmail"`
		OrganisationPhone     string `db:"OrganisationPhone"`
	}
	if err := db.Collection("configuration").Find("id", "1").One(&configuration); err != nil {
		log.Fatal(err)
	}

	config := map[string]string{
		"AdminDomain":           configuration.AdminDomain,
		"domain":                configuration.Domain,
		"OrganisationName":      configuration.OrganisationName,
		"OrganisationShortName": configuration.OrganisationShortName,
		"OrganisationEmail":     configuration.OrganisationEmail,
		"OrganisationPhone":     configuration.OrganisationPhone,
	}

	// Webserver
	r := mux.NewRouter()

	routes(r, config, db)

	srv := &http.Server{
		Handler: r,
		Addr:    ":8000",
		// Good practice: enforce timeouts for servers you create!
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Println("Mattrax Started Listening At ") //TODO: Finish This Message
	err = srv.ListenAndServe()
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
