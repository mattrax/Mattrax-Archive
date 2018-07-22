/**
 * Mattrax: An Open Source Device Management System
 * File Description: This is The Apple MDM Core. It Manages The Webserver Routes and APNS for Apples MDM.
 * Protcol Documentation: https://developer.apple.com/library/archive/documentation/Miscellaneous/Reference/MobileDeviceManagementProtocolRef/3-MDM_Protocol/MDM_Protocol.html
 * Copyright (C) 2018-2018 Oscar Beaumont <oscartbeaumont@gmail.com>
 */

package appleMDM

import (
	"fmt"
	"net/http"

	//External Deps
	"github.com/gorilla/mux" //HTTP Router

	// Internal Functions
	mlg "github.com/mattrax/Mattrax/internal/logging" //Mattrax Logging
	mcf "github.com/mattrax/Mattrax/internal/configuration" //Mattrax Configuration
	mdb "github.com/mattrax/Mattrax/internal/database" //Mattrax Database
	errors "github.com/mattrax/Mattrax/internal/errors" // Mattrax Error Handling

	// Internal Modules
	restAPI "github.com/mattrax/Mattrax/modules/appleMDM/api" //The Apple MDM REST API
)

var ( // Get The Internal State
	pgdb = mdb.GetDatabase()
	log = mlg.GetLogger()
	config = mcf.GetConfig()
)






//TODO
func Init() { log.Info("Loaded The Apple MDM Module") }

//TODO
func Mount(r *mux.Router) {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Apple Mobile Device Management Server!") }).Methods("GET")
	r.HandleFunc("/enroll", func(w http.ResponseWriter, r *http.Request) { w.Header().Set("Content-Type", "application/x-apple-aspen-config"); http.ServeFile(w, r, "data/enroll.mobileconfig") }).Methods("GET")

	//REST API
	restAPI.Mount(r.PathPrefix("/api/").Subrouter())

	// MDM Device Endpoints
	r.Handle("/inform", errors.Handler(informHandler)).Methods("PUT").HeadersRegexp("Content-Type", "application/x-apple-aspen-mdm-checkin")
	r.Handle("/server", errors.Handler(serverHandler)).Methods("PUT").HeadersRegexp("Content-Type", "application/x-apple-aspen-mdm")
}


//TODO: handle Error Even If Status Code Is 200 For errors subpackage
//TODO: Come Up With a Better name For The /server route and the /checkin route and set it up
//TODO: Reformat Apple MDM Code Files (MDM Routes Both In Independant Files)
//TODO: Redo All Log Messages, Error Codes Before v1 Release
//TODO: Handle Shutting Down In The Middle of An Inventory/Update And The The New Update Being Different
//TODO: System For Updater To Backup The DB, Install Update Restart And Do Checks And Monitoring For An Hour (Auto Rollback On Certain Errors) Before Returning To Normal
//TODO: Way To Using Postgres Get The Devices With A Policy
//Handle Devices DOSing The Server When it Keeps Failing -> Prevent It Fast And Alert The Admin

//TODO Maybe Do "/connect" Route Like The MDM Spec Wants For Trust Certs
//Read Full MDM Spec And Make Everything Bar School Manager Intergration Work
// VPP Support Through The API (Act As Proxy and Rate Limit, etc)
// Label All Functions In The appleMDM/structs/database.go File
//TODO: Is There A Need For A APNS Package?
//Software Update Caching Server Built In (Separate Module)
//TODO: Chnage /server to /checkin and /checkin to /register or something else
// Features:
//	 Auto Generate Profiles From Config Details Parsed In (Cache In Subdirectory)
//	 After Enrolling Do Inventory -> Get All The Devices Profiles, Apps, Details, Configuration, etc
//	 Detect Via APNS (And Logging) If Device Was Removed Without CheckOut Working -> Alert Admin
//	 Auto Prune Devices That Did Not Complete Enrollment -> Alert Admin
//	 Postgres Handle Database Lossing Connection and Stress Testing -> Handle Errors
//	 Clean APNS. Combine Multiple APNS Into One Request (For Bulk Without DOSing Apple)
//	 Prevent Forging Enrollment Certificates
//	 Prestage Enrollment (Template For Devices) -> DEP Support

//  See What Checkin Does If None Of The Core Values (4 Of Them) Are Not Given By The Client Does It Plist Parsing Error?
