package workflow

import (
	"testing"

	"github.com/AutOpsProject/AutOps-API/internal/domain/common"
)

func TestWorkflow(t *testing.T) {
	name := "valid-name"
	desc := "some desc"
	sourcePath := "path/to/file.zip"
	_, err := NewWorkflow("autops::project:ABCDEFGHIJ", "invalid name", desc, "invalid source path")
	if err != common.ErrInvalidName {
		t.Error("expected err to be ErrInvalidName")
	}
	_, err = NewWorkflow("autops::project:ABCDEFGHIJ", name, desc, "invalid source path")
	if err != common.ErrInvalidPathOrUrl {
		t.Error("expected err to be ErrInvalidPathOrUrl")
	}
	workflow, err := NewWorkflow("autops::project:ABCDEFGHIJ", name, desc, sourcePath)
	if err != nil {
		t.Error("expected err to be nil")
	} else if workflow.GetSourcePath() != sourcePath {
		t.Errorf("expected %s to be %s", workflow.GetSourcePath(), sourcePath)
	}
}

func TestWorkflowInputsOutputs(t *testing.T) {
	workflow, err := NewWorkflow("autops::project:ABCDEFGHIJ", "valid-name", "", "/path/to/file.zip")
	if err != nil {
		t.Error("expected err to be nil")
	}

	if len(workflow.ListInputs()) != 0 {
		t.Errorf("expected 0, got %d", len(workflow.ListInputs()))
	}
	attribute, _ := NewWorkflowAttribute(workflow.GetIdentifier().ToString(), "attributeA", "", STRING, "default")
	workflow.AddInput(attribute)
	if len(workflow.ListInputs()) != 1 {
		t.Errorf("expected 1, got %d", len(workflow.ListInputs()))
	}
	err = workflow.AddInput(attribute)
	if err != ErrWorkflowInputAlreadyPresent {
		t.Error("expected err to be ErrWorkflowInputAlreadyPresent")
	}

	attribute2, _ := NewWorkflowAttribute(workflow.GetIdentifier().ToString(), "attributeB", "", STRING, "default")
	workflow.AddInput(attribute2)
	if len(workflow.ListInputs()) != 2 {
		t.Errorf("expected 1, got %d", len(workflow.ListInputs()))
	}

	err = workflow.RemoveInput(attribute2.GetIdentifier().ToString())
	if err != nil {
		t.Error("expected err to be nil")
	}

	err = workflow.RemoveInput(attribute2.GetIdentifier().ToString())
	if err != ErrWorkflowInputNotFound {
		t.Error("expected err to be ErrWorkflowInputNotFound")
	}

	if len(workflow.ListOutputs()) != 0 {
		t.Errorf("expected 0, got %d", len(workflow.ListOutputs()))
	}
	workflow.AddOutput(attribute)
	if len(workflow.ListOutputs()) != 1 {
		t.Errorf("expected 1, got %d", len(workflow.ListOutputs()))
	}
	err = workflow.AddOutput(attribute)
	if err != ErrWorkflowOutputAlreadyPresent {
		t.Error("expected err to be ErrWorkflowOutputAlreadyPresent")
	}
	workflow.AddOutput(attribute2)
	if len(workflow.ListOutputs()) != 2 {
		t.Errorf("expected 2, got %d", len(workflow.ListOutputs()))
	}
	workflow.RemoveOutput(attribute.GetIdentifier().ToString())
	if len(workflow.ListOutputs()) != 1 {
		t.Errorf("expected 1, got %d", len(workflow.ListOutputs()))
	}
	err = workflow.RemoveOutput(attribute.GetIdentifier().ToString())
	if err != ErrWorkflowOutputNotFound {
		t.Error("expected err to be ErrWorkflowOutputNotFound")
	}
}

func TestWorkflowSteps(t *testing.T) {
	workflow, err := NewWorkflow("autops::project:ABCDEFGHIJ", "valid-name", "", "/path/to/file.zip")
	if err != nil {
		t.Error("expected err to be nil")
	}

	if len(workflow.ListSteps()) != 0 {
		t.Errorf("expected 0, got %d", len(workflow.ListSteps()))
	}
	step, _ := NewWorkflowStep(workflow.GetIdentifier().ToString(), "stepA", "", 1, nil)
	workflow.AddStep(step)
	if len(workflow.ListSteps()) != 1 {
		t.Errorf("expected 1, got %d", len(workflow.ListSteps()))
	}

	step2, _ := NewWorkflowStep(workflow.GetIdentifier().ToString(), "stepB", "", 2, nil)
	workflow.AddStep(step2)
	if len(workflow.ListSteps()) != 2 {
		t.Errorf("expected 2, got %d", len(workflow.ListSteps()))
	}

	err = workflow.RemoveStep(3)
	if err != ErrWorkflowStepNotFound {
		t.Error("expected err to be ErrWorkflowStepNotFound")
	}

	workflow.RemoveStep(1)
	if len(workflow.ListSteps()) != 1 {
		t.Errorf("expected 1, got %d", len(workflow.ListSteps()))
	}

	if workflow.ListSteps()[0].GetName() != "stepB" && workflow.ListSteps()[0].GetStepNumber() != 1 {
		t.Error("expected StepB to become step number 1 after the remove")
	}
}

func TestWorkflowRuns(t *testing.T) {
	workflow, err := NewWorkflow("autops::project:ABCDEFGHIJ", "valid-name", "", "/path/to/file.zip")
	if err != nil {
		t.Error("expected err to be nil")
	}

	if len(workflow.ListRuns()) != 0 {
		t.Errorf("expected 0, got %d", len(workflow.ListRuns()))
	}

	run, _ := NewWorkflowRun(workflow.GetIdentifier().ToString(), "runA", "")
	workflow.AddRun(run)
	if len(workflow.ListRuns()) != 1 {
		t.Errorf("expected 1, got %d", len(workflow.ListRuns()))
	}

	err = workflow.AddRun(run)
	if err != ErrWorkflowRunAlreadyPresent {
		t.Error("expected err to be ErrWorkflowRunAlreadyPresent")
	}
}
