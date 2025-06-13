package workflow

import (
	"github.com/AutOpsProject/AutOps-API/internal/domain/common"
	"github.com/AutOpsProject/AutOps-API/internal/domain/template"
)

// WorkflowStep represents a single step in a workflow.
// Each step has a unique identifier, a name, a description, a step number, and is associated with a task (template).
type WorkflowStep struct {
	common.NamedEntity
	stepNumber int
	task       *template.Template
}

// NewWorkflowStep creates a new WorkflowStep with a generated identifier.
// Returns an error if the name or description is invalid.
func NewWorkflowStep(workflowId string, name string, description string, stepNumber int, task *template.Template) (*WorkflowStep, error) {
	identifier, err := common.BuildAttributeIdentifier(workflowId, "step")
	if err != nil {
		return nil, err
	}
	return ExistingWorkflowStep(identifier.ToString(), name, description, stepNumber, task)
}

// ExistingWorkflowStep creates a WorkflowStep with the provided identifier, name, description, step number, and task.
// Returns an error if the name or description is invalid.
func ExistingWorkflowStep(identifier string, name string, description string, stepNumber int, task *template.Template) (*WorkflowStep, error) {
	namedEntity, err := common.NewNamedEntity(identifier, name, description)
	if err != nil {
		return nil, err
	}

	return &WorkflowStep{
		NamedEntity: *namedEntity,
		stepNumber:  stepNumber,
		task:        task,
	}, nil
}

// GetStepNumber returns the step number of the WorkflowStep.
func (s *WorkflowStep) GetStepNumber() int {
	return s.stepNumber
}

// GetTask returns the associated template task of the WorkflowStep.
func (s *WorkflowStep) GetTask() *template.Template {
	return s.task
}

// WorkflowStepComparator provides comparison logic between two WorkflowSteps based on their step numbers.
type WorkflowStepComparator struct{}

// Compare compares two WorkflowStep instances by their step numbers.
// Returns -1 if a < b, 1 if a > b, and 0 if they are equal.
func (WorkflowStepComparator) Compare(a *WorkflowStep, b *WorkflowStep) int {
	if a.GetStepNumber() < b.GetStepNumber() {
		return -1
	}
	if a.GetStepNumber() > b.GetStepNumber() {
		return 1
	}
	return 0
}
