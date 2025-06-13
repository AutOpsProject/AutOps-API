package policy

import (
	"testing"

	"github.com/AutOpsProject/AutOps-API/internal/domain/common"
)

func TestProjectPolicyAction(t *testing.T) {
	action := READ_PROJECT
	if action.ResourceType() != common.PROJECT {
		t.Errorf("expected %d, got %d", common.PROJECT, action.ResourceType())
	}

	str, err := action.ToString()
	if err != nil {
		t.Error("expected err to be nil")
	}
	if str != "Read" {
		t.Errorf("expected 'Read', got '%s'", str)
	}

	action = UPDATE_PROJECT
	str, err = action.ToString()
	if err != nil {
		t.Error("expected err to be nil")
	}
	if str != "Update" {
		t.Errorf("expected 'Update', got '%s'", str)
	}

	action = DELETE_PROJECT
	str, err = action.ToString()
	if err != nil {
		t.Error("expected err to be nil")
	}
	if str != "Delete" {
		t.Errorf("expected 'Delete', got '%s'", str)
	}

	action = LIST_TEMPLATES
	str, err = action.ToString()
	if err != nil {
		t.Error("expected err to be nil")
	}
	if str != "ListTemplates" {
		t.Errorf("expected 'ListTemplates', got '%s'", str)
	}

	action = LIST_WORKFLOWS
	str, err = action.ToString()
	if err != nil {
		t.Error("expected err to be nil")
	}
	if str != "ListWorkflows" {
		t.Errorf("expected 'ListWorkflows', got '%s'", str)
	}

	action = 999
	_, err = action.ToString()
	if err != ErrInvalidPolicyAction {
		t.Error("expected err to be ErrInvalidPolicyAction")
	}
}

func TestParseProjectPolicyAction(t *testing.T) {
	action, err := ParseProjectPolicyAction("Read")
	if err != nil {
		t.Errorf("expected err to be nil")
	}
	if action != READ_PROJECT {
		t.Errorf("expected %d, got %d", READ_PROJECT, action)
	}

	action, err = ParseProjectPolicyAction("Update")
	if err != nil {
		t.Errorf("expected err to be nil")
	}
	if action != UPDATE_PROJECT {
		t.Errorf("expected %d, got %d", UPDATE_PROJECT, action)
	}

	action, err = ParseProjectPolicyAction("Delete")
	if err != nil {
		t.Errorf("expected err to be nil")
	}
	if action != DELETE_PROJECT {
		t.Errorf("expected %d, got %d", DELETE_PROJECT, action)
	}

	action, err = ParseProjectPolicyAction("ListTemplates")
	if err != nil {
		t.Errorf("expected err to be nil")
	}
	if action != LIST_TEMPLATES {
		t.Errorf("expected %d, got %d", LIST_TEMPLATES, action)
	}

	action, err = ParseProjectPolicyAction("ListWorkflows")
	if err != nil {
		t.Errorf("expected err to be nil")
	}
	if action != LIST_WORKFLOWS {
		t.Errorf("expected %d, got %d", LIST_WORKFLOWS, action)
	}

	action, err = ParseProjectPolicyAction("SomethingElse")
	if err != ErrInvalidPolicyAction {
		t.Errorf("expected err to be ErrInvalidPolicyAction")
	}
	if action != -1 {
		t.Errorf("expected -1, got %d", action)
	}
}
