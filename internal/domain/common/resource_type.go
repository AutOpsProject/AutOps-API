package common

import "strings"

// ResourceType represents a type of resource in the system.
type ResourceType int

const (
	// PROJECT represents a user resource.
	USER ResourceType = iota
	// PROJECT represents a project resource.
	PROJECT
	// WORKFLOW represents a workflow resource.
	WORKFLOW
	// TEMPLATE represents a template resource.
	TEMPLATE
	// POLICY represents a policy resource.
	POLICY
)

// ToString converts a ResourceType to its string representation.
func (t ResourceType) ToString() (string, error) {
	switch t {
	case USER:
		return "user", nil
	case PROJECT:
		return "project", nil
	case WORKFLOW:
		return "workflow", nil
	case TEMPLATE:
		return "template", nil
	case POLICY:
		return "policy", nil
	default:
		return "", ErrInvalidResourceType
	}
}

// ParseResourceType parses a string into a corresponding ResourceType.
func ParseResourceType(str string) (ResourceType, error) {
	str = strings.ToLower(str)
	switch str {
	case "user":
		return USER, nil
	case "project":
		return PROJECT, nil
	case "workflow":
		return WORKFLOW, nil
	case "template":
		return TEMPLATE, nil
	case "policy":
		return POLICY, nil
	default:
		return -1, ErrInvalidResourceType
	}
}
