package errorHandling

import (
	"net/http"
	"strings"

	// External Deps
	"github.com/go-pg/pg" // Database (Postgres)

	// Internal Functions
	mcf "github.com/mattrax/mattrax/internal/configuration" //Mattrax Configuration
	mdb "github.com/mattrax/mattrax/internal/database"      //Mattrax Database
	mlg "github.com/mattrax/mattrax/internal/logging"       //Mattrax Logging
)

var pgdb = mdb.GetDatabase()
var log = mlg.GetLogger()
var config = mcf.GetConfig() // Get The Internal State
//var ( log *logrus.Logger; pgdb *pg.DB )
//func Init(_pgdb *pg.DB, _log *logrus.Logger) { pgdb = _pgdb; log = _log }

//FIXME: Redo This File. It Is A Mess And Uses Sketchy Code

func New(_msg string) error {
	return &internalError{
		errorCode: 0,
		message:   "Internal Error: " + _msg,
		fatal:     false,
	}
}

type internalError struct {
	// Code Line The Error Occured On/Was Created From
	errorCode int
	message   string
	fatal     bool //Should App Kill Everything
}

func (e *internalError) Error() string { //TODO: Understand This
	return e.message
}

func PgError(_err error) bool {
	if _err == pg.ErrNoRows || _err == pg.ErrMultiRows {
		return false
	} else {
		return true
	}
}

/*
type ErrorPG struct {
	s string
}

func (err ErrorPG) Error() string {
	return err.s
}
*/

// HTTP Error Handling
type Handler func(http.ResponseWriter, *http.Request) (int, error)

func (fn Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	returnStatus, err := fn(w, r) //returnStatus    // TODO HERE Inject: pgdb *pg.DB

	if returnStatus == 200 {
		return
	}
	if err == nil {
		return
	}

	/*
	  log.WithFields(logrus.Fields{
	    "animal": "walrus",
	    "size":   10,
	  }).Info("A group of walrus emerges from the ocean")
	*/

	switch err.(type) {
	case *internalError:
		log.Println(err)
	default:
		errorTXT := err.Error()

		if strings.HasPrefix(errorTXT, "pg:") { //TODO: This Needs To Go
			if err == pg.ErrNoRows || err == pg.ErrMultiRows {
				log.Debug("Blank PG Database")
				//} else if strings.HasPrefix(errorTXT, "pg: Model(non-pointer") { //TODO: This Need To Go Even More. IT IS BAD CODE!!
				//  log.Println("Entry Already Exists")
			} else {
				log.Println(err)
			}
		} else {
			log.Println("External Error: ", err)
		}
	}

	http.Error(w, "An Error Occured", returnStatus)

	//Is It A Database Error
	//Is It An Internal Error
	//Else

	//Handle Different Postgress Errors With Logging And Make Outright Failing And Make Everything Return Error

	/*if err != pg.ErrNoRows && err != pg.ErrMultiRows {
	  log.Warning("Postgres Error: ", err);
	   //TODO: Try Database Request Again Here
	}*/

	// If http.StatusNotFound is True Then Show Custom Error Page

	/*if returnStatus, err = fn(w, r); err != nil {
	  log.Println("HTTPS Error", err.Error())
	  http.Error(w, "A Server Side Error Occured", 500)
	}*/

	//log.Println(pg.Error)
	//log.Println(err.isNetworkError())

	//_, ok := err.(net.Error)
	//log.Println(ok)

	/*if _, ok := err.(ErrorPG); !ok {
	    log.Println("PG Error")
	}*/

	// plist Parser Error
	//    Error
	//    Parsing/Encoding Failed
	// Daatabse Errors
	//    Network
	//    Breaking Ingrity Checks
	//    etc

	//log.Println(returnStatus) //Return This HTTP Code
	//http.Error(w, "A Server Side Error Occured", 200)
}

//Handle Logging From Here. Custom Formatting Functions Probally
