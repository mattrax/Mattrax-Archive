package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	mgo "gopkg.in/mgo.v2"

	"github.com/go-kit/kit/log"
	"github.com/mattrax/Mattrax/internal/mattrax"
	"github.com/mattrax/Mattrax/internal/mongo"
	"github.com/mattrax/Mattrax/internal/server"
	appleMDM "github.com/mattrax/Mattrax/platform/apple"
	windowsMDM "github.com/mattrax/Mattrax/platform/windows"
)

//TODO: Check That No stdlib/"log" package exists

func main() {
	// TODO: Load The Config

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr)) // TODO: Can I use a custom logger
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	// Setup repositories
	var (
		devices  mattrax.DeviceRepository
		policies mattrax.PolicyRepository
		users    mattrax.UserRepository
	)

	session, err := mgo.Dial("127.0.0.1") // TODO: Load from config, setup schema/collections
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	devices, _ = mongo.NewDeviceRepository("mattrax", session) //TODO: Handle Errors
	//policies = postgres.NewPolicyRepository(db)
	//users = postgres.NewUserRepository(db)

	var as appleMDM.Service
	as = appleMDM.NewService(devices, policies, users)

	var ws windowsMDM.Service
	ws = windowsMDM.NewService(devices, policies, users)
	//TODO: Mount Logging, Circuit Breakers, Metrics To The Service

	srv := server.New(as, ws, log.With(logger, "component", "http"))

	errs := make(chan error, 2)
	go func() {
		logger.Log("transport", "http", "address", ":8000", "msg", "listening")
		errs <- http.ListenAndServe(":8000", srv)
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)
}
