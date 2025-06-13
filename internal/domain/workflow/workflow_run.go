package workflow

import (
	"strings"

	"github.com/AutOpsProject/AutOps-API/internal/domain/common"
)

// WorkflowRun represents a single execution instance of a workflow.
// It inherits identification, name, and description from NamedEntity.
type WorkflowRun struct {
	common.NamedEntity
}

// NewWorkflowRun creates a new WorkflowRun with a generated unique identifier.
// Returns an error if the name or description is invalid.
func NewWorkflowRun(workflowId string, name string, description string) (*WorkflowRun, error) {
	identifier, err := common.BuildAttributeIdentifier(workflowId, "run")
	if err != nil {
		return nil, err
	}
	return ExistingWorkflowRun(identifier.ToString(), name, description)
}

// ExistingWorkflowRun creates a WorkflowRun with the provided identifier, name, and description.
// Returns an error if the name or description is invalid.
func ExistingWorkflowRun(identifier string, name string, description string) (*WorkflowRun, error) {
	namedEntity, err := common.NewNamedEntity(identifier, name, description)
	if err != nil {
		return nil, err
	}

	return &WorkflowRun{
		NamedEntity: *namedEntity,
	}, nil
}

// WorkflowRunComparator is used to compare two WorkflowRun instances
// based on their identifier. It enables deterministic sorting within lists.
type WorkflowRunComparator struct{}

// Compare returns a comparison between two WorkflowRun identifiers.
// It is used to order attributes consistently within lists.
func (WorkflowRunComparator) Compare(a *WorkflowRun, b *WorkflowRun) int {
	return strings.Compare(a.GetIdentifier().ToString(), b.GetIdentifier().ToString())
}
