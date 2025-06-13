// Package common provides shared types and utility functions used across the domain layer.
// This file defines the Identifier type and functions for building and validating resource identifiers.
package common

import (
	"fmt"
	"regexp"
	"strings"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

// AUTOPS_ID_PREFIX is the prefix used for all AutOps identifiers.
const AUTOPS_ID_PREFIX = "autops::"

// NANO_ID_LENGTH defines the length of the NanoID used in resource identifiers.
const NANO_ID_LENGTH = 10

// Identifier represents a structured identifier for domain resources.
type Identifier struct {
	id string
}

type IdentifierComparator struct{}

func (IdentifierComparator) Compare(a, b *Identifier) int {
	return strings.Compare(a.id, b.id)
}

// ValidateIdentifier checks if the given identifier string conforms to the expected format.
// It ensures the identifier starts with the correct prefix, has a valid segment structure,
// and that each ID segment is a valid NanoID.
func ValidateIdentifier(str string) error {
	if !strings.HasPrefix(str, AUTOPS_ID_PREFIX) {
		return ErrInvalidIdentifierFormat
	}
	trimmed := strings.TrimPrefix(str, AUTOPS_ID_PREFIX)
	print(trimmed)
	segments := strings.Split(trimmed, ":")
	if len(segments) != 2 && len(segments) != 4 && len(segments) != 6 {
		return ErrInvalidIdentifierFormat
	}

	prefix, err := ParseResourceType(segments[0])
	if err != nil {
		return err
	}
	if prefix != PROJECT && (prefix != USER && len(segments) > 2) {
		return ErrInvalidIdentifierFormat
	}

	if len(segments) == 4 {
		resourceType, err := ParseResourceType(segments[2])
		if err != nil {
			return err
		}
		if resourceType == PROJECT || resourceType == USER {
			return ErrInvalidIdentifierFormat
		}
	}

	for i := 1; i < len(segments); i += 2 {
		if !isValidNanoID(segments[i]) {
			return ErrInvalidIdentifierFormat
		}
	}

	return nil
}

// isValidNanoID returns true if the given string is a valid NanoID of the expected length and character set.
func isValidNanoID(s string) bool {
	return len(s) == NANO_ID_LENGTH && regexp.MustCompile(`^[a-zA-Z0-9_-]+$`).MatchString(s)
}

// NewIdentifier creates a new Identifier instance after validating its format.
func NewIdentifier(id string) (*Identifier, error) {
	err := ValidateIdentifier(id)
	if err != nil {
		return nil, err
	}
	return &Identifier{
		id: id,
	}, nil
}

// ToString returns the string representation of the Identifier.
func (i *Identifier) ToString() string {
	return i.id
}

// Segments returns the segments of the identifier split by ':'.
func (i *Identifier) Segments() []string {
	return strings.Split(i.id, ":")
}

// GetType returns the resource type of the Identifier based on its last type segment.
func (i *Identifier) GetType() ResourceType {
	segments := i.Segments()
	t, _ := ParseResourceType(segments[len(segments)-2])
	return t
}

// GenerateNanoID generates a random NanoID of fixed length.
func GenerateNanoID() (string, error) {
	return gonanoid.New(NANO_ID_LENGTH)
}

// BuildIdentifier creates a new Identifier for a given prefix and type (<prefix>:<type>:<nano-id>)
func BuildIdentifier(prefix string, type_name string) (*Identifier, error) {
	id, err := GenerateNanoID()
	if err != nil {
		return nil, err
	}
	return NewIdentifier(fmt.Sprintf("%s:%s:%s", prefix, type_name, id))
}

// BuildUserIdentifier creates a new Identifier for a user.
func BuildUserIdentifier() (*Identifier, error) {
	return BuildIdentifier("autops:", "user")
}

// BuildProjectIdentifier creates a new Identifier for a project resource.
func BuildProjectIdentifier() (*Identifier, error) {
	return BuildIdentifier("autops:", "project")
}

// BuildTemplateIdentifier creates a new Identifier for a template under the given project.
func BuildTemplateIdentifier(projectId string) (*Identifier, error) {
	return BuildIdentifier(projectId, "template")
}

// BuildWorkflowIdentifier creates a new Identifier for a workflow under the given project.
func BuildWorkflowIdentifier(projectId string) (*Identifier, error) {
	return BuildIdentifier(projectId, "workflow")
}

// BuildPolicyIdentifier creates a new Identifier for a policy under the given project.
func BuildPolicyIdentifier(projectId string) (*Identifier, error) {
	return BuildIdentifier(projectId, "policy")
}

// BuildAttributeIdentifier creates a new Identifier for an attribute under the given parent resource and attribute type.
func BuildAttributeIdentifier(parentId string, attributeType string) (*Identifier, error) {
	return BuildIdentifier(parentId, attributeType)
}
