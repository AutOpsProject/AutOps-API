package common

import (
	"testing"
)

func TestResourceType_ToString(t *testing.T) {
	tests := []struct {
		name     string
		input    ResourceType
		expected string
		isError  bool
	}{
		{"Project", PROJECT, "project", false},
		{"Workflow", WORKFLOW, "workflow", false},
		{"Template", TEMPLATE, "template", false},
		{"Policy", POLICY, "policy", false},
		{"Invalid", ResourceType(100), "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.input.ToString()
			if tt.isError {
				if err == nil {
					t.Errorf("expected an error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result != tt.expected {
					t.Errorf("expected %v, got %v", tt.expected, result)
				}
			}
		})
	}
}

func TestParseResourceType(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected ResourceType
		isError  bool
	}{
		{"ParseProject", "project", PROJECT, false},
		{"ParseWorkflow", "workflow", WORKFLOW, false},
		{"ParseTemplate", "template", TEMPLATE, false},
		{"ParsePolicy", "policy", POLICY, false},
		{"ParseInvalid", "invalid", -1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseResourceType(tt.input)
			if tt.isError {
				if err == nil {
					t.Errorf("expected an error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result != tt.expected {
					t.Errorf("expected %v, got %v", tt.expected, result)
				}
			}
		})
	}
}
