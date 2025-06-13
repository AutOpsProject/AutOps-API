// Package policy provides types and functions for managing project policy actions.
package policy

import "github.com/AutOpsProject/AutOps-API/internal/domain/common"

// ProjectPolicyAction represents a policy action that can be performed on a project.
type ProjectPolicyAction int

const (
	// READ_PROJECT represents the action of reading a project's details.
	READ_PROJECT ProjectPolicyAction = iota
	// UPDATE_PROJECT represents the action of updating a project's information.
	UPDATE_PROJECT
	// DELETE_PROJECT represents the action of deleting a project.
	DELETE_PROJECT
	// LIST_WORKFLOWS represents the action of listing all workflows within a project.
	LIST_WORKFLOWS
	// LIST_TEMPLATES represents the action of listing all templates within a project.
	LIST_TEMPLATES
)

// ToString converts a ProjectPolicyAction to its string representation.
func (p ProjectPolicyAction) ToString() (string, error) {
	switch p {
	case READ_PROJECT:
		return "Read", nil
	case UPDATE_PROJECT:
		return "Update", nil
	case DELETE_PROJECT:
		return "Delete", nil
	case LIST_WORKFLOWS:
		return "ListWorkflows", nil
	case LIST_TEMPLATES:
		return "ListTemplates", nil
	default:
		return "", ErrInvalidPolicyAction
	}
}

// ResourceType returns the ResourceType associated with ProjectPolicyAction, which is always PROJECT.
func (p ProjectPolicyAction) ResourceType() common.ResourceType {
	return common.PROJECT
}

// ParseProjectPolicyAction parses a string into a corresponding ProjectPolicyAction.
func ParseProjectPolicyAction(action string) (ProjectPolicyAction, error) {
	switch action {
	case "Read":
		return READ_PROJECT, nil
	case "Update":
		return UPDATE_PROJECT, nil
	case "Delete":
		return DELETE_PROJECT, nil
	case "ListWorkflows":
		return LIST_WORKFLOWS, nil
	case "ListTemplates":
		return LIST_TEMPLATES, nil
	default:
		return -1, ErrInvalidPolicyAction
	}
}
