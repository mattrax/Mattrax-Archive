package mattrax

import (
	"database/sql/driver"
	"errors"
)

// ApplicationID is a unique identifier for an application.
type ApplicationID string

func (i ApplicationID) Value() (driver.Value, error) {
	return "('" + string(i) + "')", nil //fmt.Sprintf("('%s')", string(i)), nil
}

// Application is a definition of a application/piece of software that can be pushed to a device
type Application struct {
	ID                 ApplicationID
	Name               string
	SupportedPlatforms []string

	// TODO: More Platform Specific Stuff Here
}

// ErrUnknownApplication is used when a application could not be found.
var ErrUnknownApplication = errors.New("unknown application")

// ApplicationRepository provides access to a application store.
type ApplicationRepository interface { //TODO: Add to the system
	Create(policy *Application) error
	Remove(policy *Application) error
	Find(ApplicationID) (*Application, error)
}
