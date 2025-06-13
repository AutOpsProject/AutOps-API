package policy

import (
	"testing"

	"github.com/AutOpsProject/AutOps-API/internal/domain/common"
)

func TestWorkflowPolicyAction(t *testing.T) {
	action := READ_WORKFLOW
	if action.ResourceType() != common.WORKFLOW {
		t.Errorf("expected %d, got %d", common.WORKFLOW, action.ResourceType())
	}
	str, err := action.ToString()
	if err != nil {
		t.Error("expected err to be nil")
	}
	if str != "Read" {
		t.Errorf("expected 'Read', got '%s'", str)
	}

	action = UPDATE_WORKFLOW
	str, err = action.ToString()
	if err != nil {
		t.Error("expected err to be nil")
	}
	if str != "Update" {
		t.Errorf("expected 'Update', got '%s'", str)
	}

	action = DELETE_WORKFLOW
	str, err = action.ToString()
	if err != nil {
		t.Error("expected err to be nil")
	}
	if str != "Delete" {
		t.Errorf("expected 'Delete', got '%s'", str)
	}

	action = RUN_WORKFLOW
	str, err = action.ToString()
	if err != nil {
		t.Error("expected err to be nil")
	}
	if str != "Run" {
		t.Errorf("expected 'Run', got '%s'", str)
	}

	action = 999
	_, err = action.ToString()
	if err != ErrInvalidPolicyAction {
		t.Error("expected err to be ErrInvalidPolicyAction")
	}
}

func TestParseWorkflowPolicyAction(t *testing.T) {
	action, err := ParseWorkflowPolicyAction("Read")
	if err != nil {
		t.Errorf("expected err to be nil")
	}
	if action != READ_WORKFLOW {
		t.Errorf("expected %d, got %d", READ_WORKFLOW, action)
	}

	action, err = ParseWorkflowPolicyAction("Update")
	if err != nil {
		t.Errorf("expected err to be nil")
	}
	if action != UPDATE_WORKFLOW {
		t.Errorf("expected %d, got %d", UPDATE_WORKFLOW, action)
	}

	action, err = ParseWorkflowPolicyAction("Delete")
	if err != nil {
		t.Errorf("expected err to be nil")
	}
	if action != DELETE_WORKFLOW {
		t.Errorf("expected %d, got %d", DELETE_WORKFLOW, action)
	}

	action, err = ParseWorkflowPolicyAction("Run")
	if err != nil {
		t.Errorf("expected err to be nil")
	}
	if action != RUN_WORKFLOW {
		t.Errorf("expected %d, got %d", RUN_WORKFLOW, action)
	}

	action, err = ParseWorkflowPolicyAction("SomethingElse")
	if err != ErrInvalidPolicyAction {
		t.Errorf("expected err to be ErrInvalidPolicyAction")
	}
	if action != -1 {
		t.Errorf("expected -1, got %d", action)
	}
}
