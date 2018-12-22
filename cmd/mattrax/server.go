package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	logging "github.com/op/go-logging"

	"github.com/mattrax/Mattrax/internal/mattrax"
	"github.com/mattrax/Mattrax/internal/postgres"
	"github.com/mattrax/Mattrax/internal/server"
	appleMDM "github.com/mattrax/Mattrax/platform/apple"
	windowsMDM "github.com/mattrax/Mattrax/platform/windows"
)

func Server() {
	// Load The Configuration

	// Setup The Logging
	console := logging.NewBackendFormatter(logging.NewLogBackend(os.Stderr, "", 0), logging.MustStringFormatter(`%{color}%{time:15:04:05.000} %{shortfile} [%{level:.4s}] ▶%{color:reset} %{message}`))
	/*if srv.Config.Debug.LogFile != "disabled" { // TODO: After Handling The Config Work Fix This
		file, err := os.OpenFile(srv.Config.Debug.LogFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
		if err != nil {
			return err
		}
		logFile := logging.NewBackendFormatter(logging.NewLogBackend(file, "", 0), logging.MustStringFormatter(`time=%{time:2006-01-02 15:04:05.000} file=%{longfile} pid=%{pid} callpath=%{callpath} level=%{level} msg=%{message}`))
		logging.SetBackend(console, logFile)
		return nil
	}*/
	logging.SetBackend(console)
	log := logging.MustGetLogger("mattrax_cmd")

	log.Info("Starting Mattrax... Created By Oscar Beaumont")

	// Setup repositories
	var (
		devices      mattrax.DeviceRepository
		policies     mattrax.PolicyRepository
		applications mattrax.ApplicationRepository
		users        mattrax.UserRepository
	)

	// TODO: Load from config
	db, err := sql.Open("postgres", "user=oscar.beaumont dbname=mattrax sslmode=disable")
	if err != nil {
		log.Fatal("Error initialising the connection to the database", err)
		return
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Error establishing a connection with the database", err)
	}
	defer db.Close() //TODO: Error Handling

	devices = postgres.NewDeviceRepository(db)
	//policies = postgres.NewPolicyRepository(db) // TODO: Get These Working + Add Tests
	//applications = postgres.NewApplicationsRepository(db)
	//users = postgres.NewUserRepository(db)

	var as appleMDM.Service
	as = appleMDM.NewService(devices, policies, applications, users)
	as = appleMDM.NewLoggingService(as)

	var ws windowsMDM.Service
	ws = windowsMDM.NewService(devices, policies, applications, users)
	ws = appleMDM.NewLoggingService(ws)

	srv := server.New(as, ws) // TEMP: logging.MustGetLogger("mattrax_http")

	errs := make(chan error, 2)
	go func() {
		log.Info("transport=http address=" + ":8000" + " msg=listening")
		errs <- http.ListenAndServe(":8000", srv) //TODO: Gracefull Shutdown
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	log.Info("terminated", <-errs)
}
