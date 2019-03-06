package builtin

import mattrax "github.com/mattrax/Mattrax/internal"

// TODO: Finish All Of This

type UserDataStore interface {
	SaveUser(user mattrax.User) error
	RetrieveUser(user mattrax.User) error
}

type dataStoreAuthService struct {
	DataStore UserDataStore
}

func (as dataStoreAuthService) VerifyLogin(username string, password string) error {
	// TODO
	return nil
}

func (as dataStoreAuthService) VerifyEmail(email string) error {
	// TODO
	return nil
}

func (as dataStoreAuthService) VerifySessionToken(sessionToken string) error {
	// TODO
	return nil
}

func NewAuthService(ds UserDataStore) mattrax.AuthService {
	return dataStoreAuthService{ds}
}
