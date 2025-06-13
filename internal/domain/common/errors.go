package common

import (
	"errors"
)

var (
	ErrInvalidName             = errors.New("name must be non-empty, with at most 128 characters, and exclusively composed of letters, numbers, hyphens (-) and underscores (_)")
	ErrInvalidDescription      = errors.New("description must be less than 512 characters")
	ErrListIndexOutOfRange     = errors.New("index out of bound")
	ErrListNilComparator       = errors.New("comparator is nil")
	ErrStatusParseError        = errors.New("the string cannot be converted to a status")
	ErrInvalidPathOrUrl        = errors.New("the provided string does not correspond to a path or a url")
	ErrInvalidResourceType     = errors.New("invalid resource type")
	ErrInvalidIdentifierFormat = errors.New("identifier must match the  following format: 'autops::project:<project-id>[:<resource-type>:<resource-id>]'")
)
