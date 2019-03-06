package main

import (
	"net/http"

	"github.com/mattrax/Mattrax/internal/middleware"
)

const interfaceBuildDir = "./interface/build"

func (m Mattrax) routes() {
	authConfig := middleware.AuthConfig{
		LoginEndpoint: "/login",
		AuthService:   m.AuthService,
	}

	m.r.HandleFunc("/login", serveFile("/index.html")).Methods("GET")
	m.r.HandleFunc("/api/v1/login", middleware.LoginAPI(authConfig)).Methods("POST")

	// End User Interface
	endUserInterfaceHandler := serveFile("/index.html")               // TODO: Require Authentication
	m.r.HandleFunc("/enroll", endUserInterfaceHandler).Methods("GET") // TODO: Require Self Enrollment Enable or Have "admin" Role

	// Admin Interface (authenticated with "admin" role required)
	adminInterfaceHandler := middleware.AuthRequireRoles(authConfig, []string{"admin"}, serveFile("/index.html"))
	m.r.HandleFunc("/", adminInterfaceHandler).Methods("GET")
	m.r.HandleFunc("/devices", adminInterfaceHandler).Methods("GET")
	m.r.HandleFunc("/settings", adminInterfaceHandler).Methods("GET")
	m.r.HandleFunc("/enroll", adminInterfaceHandler).Methods("GET")

	// Interface Assets
	m.r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(interfaceBuildDir+"/static/")))).Methods("GET") // TODO: Caching

	// Windows MDM
	m.WindowsMDM.Routes(m.r)

	// Apple MDM
	// TODO: Apple

	// NotFound and MethodNotAllowed Handlers
	//m.r.NotFoundHandler =
	//m.r.MethodNotAllowedHandler =
}

func serveFile(file string) http.HandlerFunc {
	url := interfaceBuildDir + file
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, url)
	}
}

// TODO: Listen Only On The Correct Methods
// TODO: HTST Middleware
// TODO: s := r.Host("www.example.com").Subrouter()
