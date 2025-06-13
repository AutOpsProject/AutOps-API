package workflow

import (
	"testing"

	"github.com/AutOpsProject/AutOps-API/internal/domain/common"
	"github.com/AutOpsProject/AutOps-API/internal/domain/template"
)

func TestWorkflowStep(t *testing.T) {
	workflowId := "autops::project:ABCDEFGHIJ:workflow:1234567890"
	name := "valid-name"
	desc := "some description"
	stepNumber := 456
	template, _ := template.NewTemplate("autops::project:ABCDEFGHIJ:template:1234567890", "my-template", "", common.SUCCESS, template.TERRAFORM, "path/to/file.zip")
	_, err := NewWorkflowStep(workflowId, "invalid name", desc, stepNumber, template)
	if err != common.ErrInvalidName {
		t.Errorf("expected err to be ErrInvalidName, got %s", err)
	}

	step, err := NewWorkflowStep(workflowId, name, desc, stepNumber, template)
	if err != nil {
		t.Error("expected err to be nil")
	} else if step.GetStepNumber() != stepNumber {
		t.Errorf("expected %d to be %d", step.GetStepNumber(), stepNumber)
	} else if step.GetTask() != template {
		t.Errorf("expected %p to be %p", step.GetTask(), template)
	}
}

func TestCompareWorkflowStep(t *testing.T) {
	workflowId := "autops::project:ABCDEFGHIJ:workflow:1234567890"
	template, _ := template.NewTemplate("autops::project:ABCDEFGHIJ:template:1234567890", "my-template", "", common.SUCCESS, template.TERRAFORM, "path/to/file.zip")
	stepA, _ := NewWorkflowStep(workflowId, "valid-name", "", 2, template)
	stepB, _ := NewWorkflowStep(workflowId, "valid-name", "", 1, template)
	comparator := WorkflowStepComparator{}

	if comparator.Compare(stepA, stepB) != 1 {
		t.Errorf("expected 1, got %d", comparator.Compare(stepA, stepB))
	}
	if comparator.Compare(stepB, stepA) != -1 {
		t.Errorf("expected -1, got %d", comparator.Compare(stepB, stepA))
	}
	if comparator.Compare(stepA, stepA) != 0 {
		t.Errorf("expected 0, got %d", comparator.Compare(stepA, stepA))
	}
}
