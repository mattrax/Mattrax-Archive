package mattrax

import (
	"github.com/gorilla/mux"
)

type MDM interface {
	Routes(*mux.Router) error
}
