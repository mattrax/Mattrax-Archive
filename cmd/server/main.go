package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

type Server struct {
	router *httprouter.Router
	log    *zap.Logger

	srv http.Server
	//db     *pg.DB
	//config Config
}

func (server *Server) SetupLogger(production bool) {
	var err error
	if production {
		server.log, err = zap.NewProduction()
		if err != nil {
			panic(err)
		}
	} else {
		server.log, err = zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
	}
}

func (server *Server) ListenHTTP(addr string) {
	server.srv = http.Server{
		Addr:    addr,
		Handler: server.router, // Handler(server.router), // handlers.LoggingHandler(os.Stdout, router), //TODO: Logging Handler (Check Disabled Production Preformance)
		//ErrorLog:     logger, //TODO: Does This Matter
		ReadTimeout:  5 * time.Second, //TODO: Config Options To Override
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		signal.Stop(quit)
		fmt.Println("Shutting Down")
		ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
		server.srv.Shutdown(ctx)
	}()

	server.log.Info("Started Mattrax At " + addr)
	log.Fatal(server.srv.ListenAndServe()) //TODO: Only Run This If Error
}

func main() {
	server := Server{
		router: httprouter.New(),
	}
	server.SetupLogger(false)
	server.routes()
	server.ListenHTTP(":8000") //TODO: Cut Down On All The Imports/Code This Uses If Possible
}

//TODO: Document Every Function/Struct

///// TEMP /////
/*
type Handler struct {
	next http.Handler
}

// Make a constructor for our middleware type since its fields are not exported (in lowercase)
func NewMiddleware(next http.Handler) *Handler {
	return &Handler{next: next}
}

// Our middleware handler
func (s *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// We can modify the request here; for simplicity, we will just log a message
	log.Printf("msg: %s, Method: %s, URI: %s\n", r.Method, r.RequestURI)
	s.next.ServeHTTP(w, r)
}*/
