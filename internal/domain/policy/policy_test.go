package policy

import (
	"testing"

	"github.com/AutOpsProject/AutOps-API/internal/domain/common"
)

func TestPolicy_ListStatements(t *testing.T) {
	id, _ := common.NewIdentifier("autops::project:1234567890")
	statement1 := &PolicyStatement{
		resourceIdentifiers: common.NewList(nil, []*common.Identifier{id}),
		actions:             common.NewList[PolicyAction](nil, asPolicyActions(&mockPolicyAction{name: "Read", rt: common.PROJECT})),
		effect:              ALLOW,
	}
	policy, err := NewPolicy("autops::project:1234567890", "TestPolicy", "desc", []*PolicyStatement{statement1})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	stmts := policy.ListStatements()
	if len(stmts) != 1 || stmts[0] != statement1 {
		t.Errorf("ListStatements() returned incorrect result")
	}
}

func TestPolicy_GetPermission(t *testing.T) {
	action := &mockPolicyAction{name: "Read", rt: common.PROJECT}
	id, _ := common.NewIdentifier("autops::project:1234567890")
	statementAllow := &PolicyStatement{
		resourceIdentifiers: common.NewList(common.IdentifierComparator{}, []*common.Identifier{id}),
		actions:             common.NewList[PolicyAction](PolicyActionComparator{}, asPolicyActions(action)),
		effect:              ALLOW,
	}

	statementDeny := &PolicyStatement{
		resourceIdentifiers: common.NewList(common.IdentifierComparator{}, []*common.Identifier{id}),
		actions:             common.NewList[PolicyAction](PolicyActionComparator{}, asPolicyActions(action)),
		effect:              DENY,
	}

	policy, _ := NewPolicy("autops::project:1234567890", "test", "test", []*PolicyStatement{})
	effect := policy.GetPermission(id, action)
	if effect != UNSPECIFIED {
		t.Errorf("expected %d, got %d", UNSPECIFIED, effect)
	}

	policy.statements.Append(statementAllow)
	effect = policy.GetPermission(id, action)
	if effect != ALLOW {
		t.Errorf("expected %d, got %d", ALLOW, effect)
	}

	policy.statements.Append(statementDeny)
	effect = policy.GetPermission(id, action)
	if effect != DENY {
		t.Errorf("expected %d, got %d", DENY, effect)
	}

	policy.statements.Remove(statementAllow)
	effect = policy.GetPermission(id, action)
	if effect != DENY {
		t.Errorf("expected %d, got %d", DENY, effect)
	}
}
