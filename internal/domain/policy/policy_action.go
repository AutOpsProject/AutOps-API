package policy

import (
	"fmt"
	"strings"

	"github.com/AutOpsProject/AutOps-API/internal/domain/common"
)

// PolicyAction represents an action that can be performed on a specific resource type.
// Each action must be convertible to a string and must provide its associated resource type.
type PolicyAction interface {
	ToString() (string, error)
	ResourceType() common.ResourceType
}

type PolicyActionComparator struct{}

func (PolicyActionComparator) Compare(a, b PolicyAction) int {
	a_str, _ := a.ToString()
	b_str, _ := b.ToString()
	return strings.Compare(a_str, b_str)
}

// GetFullName returns the full name of a policy action in the format "action:resource_type".
// It combines the action name and its resource type into a single identifier string.
func GetFullName(p PolicyAction) (string, error) {
	str, err := p.ToString()
	if err != nil {
		return "", err
	}
	rt, err := p.ResourceType().ToString()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%s", str, rt), nil
}

// ParsePolicyAction parses a string formatted as "resource_type:action" and returns the corresponding PolicyAction.
// It supports parsing actions for known resource types such as "project" and "workflow".
func ParsePolicyAction(str string) (PolicyAction, error) {
	parts := strings.SplitN(str, ":", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid action format: %s", str)
	}
	resource, action := parts[0], parts[1]

	switch resource {
	case "project":
		return ParseProjectPolicyAction(action)
	case "workflow":
		return ParseWorkflowPolicyAction(action)
	default:
		return nil, fmt.Errorf("unknown resource type: %s", resource)
	}
}
