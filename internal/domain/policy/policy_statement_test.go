package policy

import (
	"testing"

	"github.com/AutOpsProject/AutOps-API/internal/domain/common"
)

func TestPolicyEffect(t *testing.T) {
	effect := ALLOW
	str, err := effect.ToString()
	if err != nil {
		t.Error("expected err to be nil")
	}
	if str != "Allow" {
		t.Errorf("expected 'Allow', got %s", str)
	}

	effect = DENY
	str, err = effect.ToString()
	if err != nil {
		t.Error("expected err to be nil")
	}
	if str != "Deny" {
		t.Errorf("expected 'Deny', got %s", str)
	}

	effect = UNSPECIFIED
	_, err = effect.ToString()
	if err != ErrInvalidPolicyEffect {
		t.Error("expected err to be ErrInvalidPolicyEffect")
	}

	effect, err = ParsePolicyEffect("Allow")
	if err != nil {
		t.Error("expected err to be nil")
	}
	if effect != ALLOW {
		t.Errorf("expected %d, got %d", ALLOW, effect)
	}

	effect, err = ParsePolicyEffect("Deny")
	if err != nil {
		t.Error("expected err to be nil")
	}
	if effect != DENY {
		t.Errorf("expected %d, got %d", DENY, effect)
	}

	effect, err = ParsePolicyEffect("SomethingElse")
	if err != ErrInvalidPolicyEffect {
		t.Error("expected err to be ErrInvalidPolicyEffect")
	}
	if effect != -1 {
		t.Errorf("expected -1, got %d", effect)
	}
}

func TestNewPolicyStatement(t *testing.T) {
	id1, _ := common.NewIdentifier("autops::project:ABCDEFGHIJ")
	id2, _ := common.NewIdentifier("autops::project:1234567890")
	statement, err := NewPolicyStatement(ALLOW, []*common.Identifier{id1, id2}, []PolicyAction{})
	if err != nil {
		t.Error("expected err to be nil")
	}
	if statement == nil {
		t.Error("expected statement to be not nil")
	}

	statement, err = NewPolicyStatement(999, []*common.Identifier{id1, id2}, []PolicyAction{})
	if err != ErrInvalidPolicyEffect {
		t.Error("expected err to be ErrInvalidPolicyEffect")
	}
	if statement != nil {
		t.Error("expected statement to be nil")
	}
}

func TestPolicyStatement_ListResources(t *testing.T) {
	res1, _ := common.NewIdentifier("autops::project:1234567890")
	res2, _ := common.NewIdentifier("autops::project:AZERTYUIOP")

	stmt, err := NewPolicyStatement(ALLOW, []*common.Identifier{res1, res2}, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := stmt.ListResources()
	if len(got) != 2 || got[0] != res1 || got[1] != res2 {
		t.Errorf("ListResources returned unexpected result: %v", got)
	}
}

func TestPolicyStatement_ListActions(t *testing.T) {
	act1 := mockPolicyAction{name: "Read"}
	act2 := mockPolicyAction{name: "Update"}

	stmt, err := NewPolicyStatement(ALLOW, nil, []PolicyAction{&act1, &act2})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := stmt.ListActions()
	if len(got) != 2 || got[0] != &act1 || got[1] != &act2 {
		t.Errorf("ListActions returned unexpected result: %v", got)
	}
}
