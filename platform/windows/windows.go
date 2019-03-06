package windows

import (
	mattrax "github.com/mattrax/Mattrax/internal"
)

type windowsMDM struct {
	DataStore   mattrax.DataStore
	AuthService mattrax.AuthService
}

func MDM(ds mattrax.DataStore, as mattrax.AuthService) mattrax.MDM {
	// TODO: Init The Datastore
	return windowsMDM{ds, as}
}
