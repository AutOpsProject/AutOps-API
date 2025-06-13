package identity

import "errors"

var (
	ErrAttachedPolicyNotFound = errors.New("cannot find a policy with the provided identifer attached to the current restricted entity")
	ErrInvalidEmail           = errors.New("the provided string does not match a valid email address")
	ErrInvalidUsername        = errors.New("username length must be 3-30 characters, and only composed of letters, number and underscores '_'")
)
