package policy

import (
	"testing"

	"github.com/AutOpsProject/AutOps-API/internal/domain/common"
)

type mockPolicyAction struct {
	name string
	rt   common.ResourceType
}

func (m *mockPolicyAction) ToString() (string, error) {
	return m.name, nil
}

func (m *mockPolicyAction) ResourceType() common.ResourceType {
	return m.rt
}

func asPolicyActions(actions ...*mockPolicyAction) []PolicyAction {
	res := make([]PolicyAction, len(actions))
	for i, a := range actions {
		res[i] = a
	}
	return res
}

func TestGetFullName(t *testing.T) {
	tests := []struct {
		name     string
		action   PolicyAction
		expected string
		wantErr  bool
	}{
		{
			name:     "Valid project action",
			action:   &mockPolicyAction{name: "Read", rt: common.PROJECT},
			expected: "Read:project",
			wantErr:  false,
		},
		{
			name:     "Invalid resource type",
			action:   &mockPolicyAction{name: "Read", rt: common.ResourceType(999)},
			expected: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFullName(tt.action)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFullName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("GetFullName() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestParsePolicyAction(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "Valid project action",
			input:   "project:Read",
			wantErr: false,
		},
		{
			name:    "Valid workflow action",
			input:   "workflow:Run",
			wantErr: false,
		},
		{
			name:    "Invalid format",
			input:   "invalidFormat",
			wantErr: true,
		},
		{
			name:    "Unknown resource",
			input:   "unknown:Something",
			wantErr: true,
		},
		{
			name:    "Unknown project action",
			input:   "project:DoSomethingStrange",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParsePolicyAction(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParsePolicyAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
