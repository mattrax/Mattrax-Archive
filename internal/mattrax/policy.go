package mattrax

import "errors"

// PolicyID is a unique identifier for the policy.
type PolicyID string

// Policy is a definition of configuration that can be pushed to a device
type Policy struct {
	ID                 PolicyID
	Name               string
	SupportedPlatforms []string

	// TODO: More Platform Specific Stuff Here
}

// ErrUnknownPolicy is used when a policy could not be found.
var ErrUnknownPolicy = errors.New("unknown Policy")

// PolicyRepository provides access to a policy store.
type PolicyRepository interface {
	Create(policy *Policy) error
	Remove(policy *Policy) error
	Find(PolicyID) (*Policy, error)
}
