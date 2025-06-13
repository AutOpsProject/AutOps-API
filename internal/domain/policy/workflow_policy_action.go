package policy

import "github.com/AutOpsProject/AutOps-API/internal/domain/common"

// WorkflowPolicyAction defines the set of possible actions that can be performed on a workflow resource.
type WorkflowPolicyAction int

const (
	READ_WORKFLOW WorkflowPolicyAction = iota
	UPDATE_WORKFLOW
	DELETE_WORKFLOW
	RUN_WORKFLOW
)

// ToString returns the string representation of a WorkflowPolicyAction.
// It returns an error if the action is not recognized.
func (p WorkflowPolicyAction) ToString() (string, error) {
	switch p {
	case READ_WORKFLOW:
		return "Read", nil
	case UPDATE_WORKFLOW:
		return "Update", nil
	case DELETE_WORKFLOW:
		return "Delete", nil
	case RUN_WORKFLOW:
		return "Run", nil
	default:
		return "", ErrInvalidPolicyAction
	}
}

// ResourceType returns the ResourceType associated with WorkflowPolicyAction, which is WORKFLOW.
func (p WorkflowPolicyAction) ResourceType() common.ResourceType {
	return common.WORKFLOW
}

// ParseWorkflowPolicyAction converts a string to a corresponding WorkflowPolicyAction.
// Returns an error if the string does not match a known action.
func ParseWorkflowPolicyAction(action string) (WorkflowPolicyAction, error) {
	switch action {
	case "Read":
		return READ_WORKFLOW, nil
	case "Update":
		return UPDATE_WORKFLOW, nil
	case "Delete":
		return DELETE_WORKFLOW, nil
	case "Run":
		return RUN_WORKFLOW, nil
	default:
		return -1, ErrInvalidPolicyAction
	}
}
