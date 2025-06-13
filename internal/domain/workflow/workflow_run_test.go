package workflow

import (
	"testing"

	"github.com/AutOpsProject/AutOps-API/internal/domain/common"
)

func TestWorkflowRun(t *testing.T) {
	workflowId := "autops::project:ABCDEFGHIJ:workflow:1234567890"
	_, err := NewWorkflowRun(workflowId, "invalid name", "some description")
	if err != common.ErrInvalidName {
		t.Error("expected err to be ErrInvalidName")
	}

	name := "valid-name"
	desc := "some description"
	workflowRun, err := NewWorkflowRun(workflowId, name, desc)
	if err != nil {
		t.Error("expected err to be nil")
	} else if workflowRun.GetName() != name {
		t.Errorf("expected %s to be %s", workflowRun.GetName(), name)
	}
}
