package mattrax

import "errors"

// UserID is a unique identifier for the user.
type UserID string

// User is an administrator or end user that can authenticate with the MDM
type User struct {
	ID       UserID
	Name     string
	Username string
	Password string

	Permissions map[string]string
}

// ErrUnknownUser is used when a user could not be found.
var ErrUnknownUser = errors.New("unknown user")

// UserRepository provides access to a user store.
type UserRepository interface {
	Create(user *User) error
	Remove(user *User) error
	Find(UserID) (*User, error)
}
